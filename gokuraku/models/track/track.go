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

type Track struct {
	Id           string `redis:"id"`
	Url          string `redis:"url"`
	Genre        string `redis:"genre"`
	UserImageUrl string `redis:"user_image_url"`
	UserUrl      string `redis:"user_url"`
	UserName     string `redis:"user_name"`
	ImageUrl     string `redis:"image_url"`
	Description  string `redis:"description"`
	CreatedAt    string `redis:"created_at"`
	Title        string `redis:"title"`
	Type         string `redis:"type"`
}

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

func CreateFromUrl(track_url string) (*Track, error) {
	var err error

	newTrack, err := NewFromUrl(track_url)

	if err != nil {
		return nil, err
	}

	err = newTrack.Save()

	if err != nil {
		return nil, err
	}

	return newTrack, nil
}

func NewFromUrl(track_url string) (*Track, error) {
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

func (t *Track) Save() error {
	var err error
	redisClient := redis.Get()
	defer redisClient.Close()

	_, err = redisClient.Do("SADD", "gokuraku:track-ids", t.Id)

	if err != nil {
		return err
	}

	err = t.Validate()

	if err != nil {
		return err
	}

	_, err = redisClient.Do(
		"HMSET", "gokuraku:track:"+t.Id,
		"id", t.Id,
		"title", t.Title,
		"type", t.Type,
		"description", t.Description,
		"image_url", t.ImageUrl,
		"user_image_url", t.UserImageUrl,
		"user_url", t.UserUrl,
		"user_name", t.UserName,
		"url", t.Url,
		"genre", t.Genre,
		"created_at", t.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func Page(page_num int) ([]*Track, error) {
	redisClient := redis.Get()
	defer redisClient.Close()

	trackIds, _ := redis.Strings(redisClient.Do("SORT", "gokuraku:track-ids", "By", "gokuraku:track:*->created_at", "LIMIT", 10*page_num, 10*(page_num+1), "ALPHA", "DESC"))

	tracks := []*Track{}
	for _, trackId := range trackIds {
		track, err := Find(trackId)

		if err != nil {
			log.Panicln(err)
		}

		tracks = append(tracks, &track)
	}

	return tracks, nil
}

func All() ([]*Track, error) {
	redisClient := redis.Get()
	defer redisClient.Close()

	trackIds, _ := redis.Strings(redisClient.Do("SORT", "gokuraku:track-ids", "By", "gokuraku:track:*->created_at", "ALPHA", "DESC"))

	tracks := []*Track{}
	for _, trackId := range trackIds {
		track, err := Find(trackId)

		if err != nil {
			log.Panicln(err)
		}

		tracks = append(tracks, &track)
	}

	return tracks, nil
}

func Find(id string) (Track, error) {
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

func (t *Track) Destroy() {
	var err error
	redisClient := redis.Get()
	defer redisClient.Close()
	_, err = redisClient.Do("SREM", "gokuraku:track-ids", t.Id)

	if err != nil {
		log.Panicln(err)
	}

	_, err = redisClient.Do("DEL", "gokuraku:track:"+t.Id)

	if err != nil {
		log.Panicln(err)
	}
}

func (t Track) Validate() error {
	if t.Id == "" {
		return errors.New("This item doesn't have ID")
	}

	if t.Url == "" {
		return errors.New("This item doesn't have Url")
	}

	if t.Title == "" {
		return errors.New("This item doesn't have Title")
	}

	if t.UserName == "" {
		return errors.New("This item doesn't have UserName")
	}

	if t.UserUrl == "" {
		return errors.New("This item doesn't have UserUrl")
	}

	return nil
}

func getRandam() *Track {
	redisClient := redis.Get()
	defer redisClient.Close()

	trackId, _ := redis.String(redisClient.Do("SRANDMEMBER", "gokuraku:track-ids"))
	values, err := redis.Values(redisClient.Do("HGETALL", "gokuraku:track:"+trackId))
	if err != nil {
		panic(err)
	}

	track := Track{}
	err = redis.ScanStruct(values, &track)
	if err != nil {
		panic(err)
	}

	return &track
}
