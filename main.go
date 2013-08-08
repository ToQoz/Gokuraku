package main

import (
	"github.com/ToQoz/Gokuraku/gokuraku/http"
	"github.com/ToQoz/Gokuraku/gokuraku/router"
	"github.com/ToQoz/Gokuraku/gokuraku/websocket"
)

func main() {
	router.Map()
	go websocket.Run()
	http.Run()
}
