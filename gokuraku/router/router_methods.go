package router

import (
	"github.com/ToQoz/Gokuraku/gokuraku"
	"github.com/ToQoz/Gokuraku/gokuraku/controllers"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"os"
)

func Map() {
	mapStatics()
	mapTracks()
}

func mapStatics() {
	root := os.Getenv("GOPATH") + "/src/github.com/ToQoz/Gokuraku"

	goweb.MapStaticFile("/", root+"/public/html/index.html")
	goweb.MapStatic("/js", root+"/public/js")
	goweb.MapStatic("/css", root+"/public/css")
	goweb.MapStaticFile("/favicon.ico", root+"/public/favicon.ico")
	goweb.Map("/soundcloud_client_id", func(ctx context.Context) error {
		return goweb.API.RespondWithData(ctx, gokuraku.Config.SoundcloudClientId)
	})
	goweb.Map("/websocket_port", func(ctx context.Context) error {
		return goweb.API.RespondWithData(ctx, gokuraku.Config.WebSocketPort)
	})
}

func mapTracks() {
	tracksController := new(controllers.TracksController)
	goweb.MapController(tracksController)
}
