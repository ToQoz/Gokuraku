package track

import (
	"errors"
	"github.com/ToQoz/Gokuraku/gokuraku/redis"
	"log"
	"strconv"
	"time"
)

type CurrentTrack struct {
	TrackId   string `redis:"track_id"`
	StartedAt string `redis:"started_at"`
	*Track
}

func (ct *CurrentTrack) Validate() error {
	if ct.TrackId == "" {
		return errors.New("Current track should have track ID")
	}

	return nil
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

	track, err := Find(currentTrack.TrackId)

	if err != nil {
		log.Panicln(err)
	}

	return &CurrentTrack{currentTrack.TrackId, currentTrack.StartedAt, &track}, nil
}

func (ct *CurrentTrack) Destroy() {
	var err error
	redisClient := redis.Get()
	defer redisClient.Close()

	_, err = redisClient.Do("DEL", "gokuraku:current-track")

	if err != nil {
		log.Panicln(err)
	}
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
