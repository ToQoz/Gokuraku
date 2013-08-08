package soundcloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	ApiEndPoint = "http://api.soundcloud.com"
)

type User struct {
	Name     string `json:"username"`
	Url      string `json:"permalink_url"`
	ImageUrl string `json:"avatar_url"`
}

type Item struct {
	Id          int    `json:"id"`
	User        User   `json:"user"`
	Genre       string `json:"genre"`
	ImageUrl    string `json:"artwork_url"`
	Type        string `json:"kind"`
	Description string `json:"description"`
	Streamable  bool   `json:"streamable"`
	Title       string `json:"title"`
	Url         string `json:"permalink_url"`
}

type API struct {
	ClientId string
}

func (api *API) Resolve(url string) (*Item, error) {
	var err error

	api_url := ApiEndPoint + "/resolve.json?url=" + url + "&client_id=" + api.ClientId

	log.Println("Access SoundCloud API: " + api_url)
	resp, err := http.Get(api_url)

	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Fail to get <%s>", api_url))
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	item := new(Item)
	err = json.NewDecoder(resp.Body).Decode(item)

	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Fail to convert json from <%s>", api_url))
	}

	return item, nil
}
