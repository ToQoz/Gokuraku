package track

import (
	"errors"
	"github.com/ToQoz/Gokuraku/gokuraku/redis"
	"log"
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

func (ct *CurrentTrack) Destroy() {
	var err error
	redisClient := redis.Get()
	defer redisClient.Close()

	_, err = redisClient.Do("DEL", "gokuraku:current-track")

	if err != nil {
		log.Panicln(err)
	}
}
