package gokuraku

import (
	"flag"
	"fmt"
)

var Config = &struct {
	HttpAddr           string
	WebSocketAddr      string
	RedisAddr          string
	RedisPassword      string
	SoundcloudClientId string
}{}

func init() {
	http_addr := flag.String("http_addr", ":9090", "http server listen address(HOST:PORT)")
	ws_addr := flag.String("ws_addr", ":9099", "websocket server listen address(HOST:PORT)")
	redis_addr := flag.String("redis_addr", ":6379", "redis address(HOST:PORT)")
	redis_password := flag.String("redis_password", "", "redis password")
	soundcloud_client_id := flag.String("soundcloud_client_id", "", "soundcloud client key(HOST:PORT)")
	flag.Parse()

	Config.HttpAddr = *http_addr
	Config.WebSocketAddr = *ws_addr
	Config.RedisAddr = *redis_addr
	Config.RedisPassword = *redis_password
	Config.SoundcloudClientId = *soundcloud_client_id

	// Disable for go test... if I come up with good idea. These will be reborn...
	// if Config.SoundcloudClientId == "" {
	// 	log.Fatalln("require soundcloud_client_id")
	// }

	fmt.Println("<Current Config>")
	fmt.Println("  * You can set your custom value by flag. see `$ gokuraku --help`")
	fmt.Printf("  gokuraku.Config.HttpAddr: %s\n", Config.HttpAddr)
	fmt.Printf("  gokuraku.Config.WebSocketAddr: %s\n", Config.WebSocketAddr)
	fmt.Printf("  gokuraku.Config.RedisAddr: %s\n", Config.RedisAddr)
	fmt.Printf("  gokuraku.Config.SoundcloudClientId: %s\n", Config.SoundcloudClientId)
	fmt.Println("</Current Config>")
	fmt.Println("")
}
