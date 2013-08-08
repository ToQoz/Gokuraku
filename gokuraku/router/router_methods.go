package router

import (
	"github.com/ToQoz/Gokuraku/gokuraku"
	"github.com/ToQoz/Gokuraku/gokuraku/controllers"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
)

func Map() {
	mapStatics()
	mapTracks()
}

func mapStatics() {
	goweb.MapStaticFile("/", "public/html/index.html")
	goweb.MapStatic("/js", "public/js")
	goweb.MapStatic("/css", "public/css")
	goweb.MapStaticFile("/favicon.ico", "public/favicon.ico")
	goweb.Map("/soundcloud_client_id", func(ctx context.Context) error {
		return goweb.API.RespondWithData(ctx, gokuraku.Config.SoundcloudClientId)
	})
}

func mapTracks() {
	tracksController := new(controllers.TracksController)
	goweb.MapController(tracksController)
}
