package track

import (
	"errors"
	"fmt"
	"github.com/ToQoz/Gokuraku/gokuraku"
	"github.com/ToQoz/Gokuraku/gokuraku/redis"
	"github.com/ToQoz/Gokuraku/gokuraku/soundcloud"
	"log"
	"net/url"
	"strconv"
	"time"
)

func TrackCount() int {
	var err error
	redisClient := redis.Get()
	defer redisClient.Close()
	count, err := redis.Int(redisClient.Do("SCARD", "gokuraku:track-ids"))

	if err != nil {
		log.Panicln(err)
	}

	return count
}

func Page(page_num int) ([]*Track, error) {
	redisClient := redis.Get()
	defer redisClient.Close()

	trackIds, _ := redis.Strings(redisClient.Do("SORT", "gokuraku:track-ids", "By", "gokuraku:track:*->created_at", "LIMIT", 10*page_num, 10*(page_num+1), "ALPHA", "DESC"))

	tracks := []*Track{}
	for _, trackId := range trackIds {
		track, err := FindTrack(trackId)

		if err != nil {
			log.Panicln(err)
		}

		tracks = append(tracks, &track)
	}

	return tracks, nil
}

func TrackAll() ([]*Track, error) {
	redisClient := redis.Get()
	defer redisClient.Close()

	trackIds, _ := redis.Strings(redisClient.Do("SORT", "gokuraku:track-ids", "By", "gokuraku:track:*->created_at", "ALPHA", "DESC"))

	tracks := []*Track{}
	for _, trackId := range trackIds {
		track, err := FindTrack(trackId)

		if err != nil {
			log.Panicln(err)
		}

		tracks = append(tracks, &track)
	}

	return tracks, nil
}

func FindTrack(id string) (Track, error) {
	redisClient := redis.Get()
	defer redisClient.Close()

	values, err := redis.Values(redisClient.Do("HGETALL", "gokuraku:track:"+id))
	if err != nil {
		log.Panicln(err)
	}

	track := Track{}
	err = redis.ScanStruct(values, &track)
	if err != nil {
		log.Panicln(err)
	}

	return track, nil
}

func CreateTrackFromUrl(track_url string) (*Track, error) {
	var err error

	newTrack, err := NewTrackFromUrl(track_url)

	if err != nil {
		return nil, err
	}

	err = newTrack.Save()

	if err != nil {
		return nil, err
	}

	return newTrack, nil
}

func NewTrackFromUrl(track_url string) (*Track, error) {
	now := strconv.FormatInt(time.Now().Unix(), 10)

	parsedUrl, parseErr := url.Parse(track_url)
	if parseErr != nil {
		return nil, parseErr
	}

	if parsedUrl.Host == "soundcloud.com" {
		api := soundcloud.API{ClientId: gokuraku.Config.SoundcloudClientId}

		item, apiErr := api.Resolve(track_url)
		if apiErr != nil {
			return nil, apiErr
		}

		if item.Type != "track" {
			return nil, errors.New("This is not track")
		}

		if item.Streamable != true {
			return nil, errors.New("This is not streamable")
		}

		track := Track{
			Id:           fmt.Sprintf("%d", item.Id),
			Genre:        item.Genre,
			Title:        item.Title,
			Description:  item.Description,
			Url:          item.Url,
			ImageUrl:     item.ImageUrl,
			UserImageUrl: item.User.ImageUrl,
			UserUrl:      item.User.Url,
			UserName:     item.User.Name,
			Type:         "soundcloud",
			CreatedAt:    now,
		}
		return &track, nil
	}

	return nil, errors.New("Unknown service " + parsedUrl.Host)
}

func Next() (*CurrentTrack, error) {
	err := updateCurrent()

	if err != nil {
		return nil, err
	}

	currentTrack, err := GetCurrent()

	if err != nil {
		return nil, err
	}

	return currentTrack, nil
}

func GetCurrent() (*CurrentTrack, error) {
	var err error
	redisClient := redis.Get()
	defer redisClient.Close()

	exists, _ := redis.Bool(redisClient.Do("EXISTS", "gokuraku:current-track"))
	if exists == false {
		err := updateCurrent()

		if err != nil {
			return nil, err
		}

		return GetCurrent()
	}

	values, err := redis.Values(redisClient.Do("HGETALL", "gokuraku:current-track"))
	currentTrack := CurrentTrack{}
	err = redis.ScanStruct(values, &currentTrack)

	if err != nil {
		log.Panicln(err)
	}

	err = currentTrack.Validate()

	if err != nil {
		currentTrack.Destroy()
		return nil, err
	}

	track, err := FindTrack(currentTrack.TrackId)

	if err != nil {
		log.Panicln(err)
	}

	return &CurrentTrack{currentTrack.TrackId, currentTrack.StartedAt, &track}, nil
}

func updateCurrent() error {
	var err error

	if TrackCount() == 0 {
		return errors.New("Track is not exists")
	}

	redisClient := redis.Get()
	defer redisClient.Close()

	t := getRandam()
	err = t.Validate()

	if err != nil {
		t.Destroy()
		return err
	}

	_, err = redisClient.Do(
		"HMSET", "gokuraku:current-track",
		"track_id", t.Id,
		"started_at", strconv.FormatInt(time.Now().Unix(), 10),
	)

	return err
}
