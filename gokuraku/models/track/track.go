package track

import (
	"errors"
	"github.com/ToQoz/Gokuraku/gokuraku/redis"
	"log"
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
