package controllers

import (
	"errors"
	"github.com/ToQoz/Gokuraku/gokuraku/models/track"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net/http"
)

type TracksController struct {
	Tracks []*track.Track
}

func (c *TracksController) Create(ctx context.Context) error {
	data, dataErr := ctx.RequestData()
	if dataErr != nil {
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, dataErr.Error())
	}
	dataMap := data.(map[string]interface{})

	if dataMap["Url"] == nil {
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, errors.New("URL required").Error())
	}

	track_url := dataMap["Url"].(string)
	track, err := track.CreateTrackFromUrl(track_url)

	if err != nil {
		log.Println(err.Error())
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	return goweb.API.Respond(ctx, http.StatusCreated, track, nil)
}

func (c *TracksController) ReadMany(ctx context.Context) error {
	tracks, _ := track.Page(0)
	return goweb.API.RespondWithData(ctx, tracks)
}
