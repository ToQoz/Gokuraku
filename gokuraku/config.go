package gokuraku

import (
	"flag"
	"fmt"
)

var Config = &struct {
	HttpPort           string
	WebSocketPort      string
	RedisAddr          string
	RedisPassword      string
	SoundcloudClientId string
}{}

func init() {
	http_port := flag.String("p", "9090", "http server listen port")
	ws_port := flag.String("ws_p", "9099", "websocket server listen port")
	redis_addr := flag.String("redis_addr", ":6379", "redis address(HOST:PORT)")
	redis_password := flag.String("redis_password", "", "redis password")
	soundcloud_client_id := flag.String("soundcloud_client_id", "", "soundcloud client key")
	flag.Parse()

	Config.HttpPort = *http_port
	Config.WebSocketPort = *ws_port
	Config.RedisAddr = *redis_addr
	Config.RedisPassword = *redis_password
	Config.SoundcloudClientId = *soundcloud_client_id

	// Disable for go test... if I come up with good idea. These will be reborn...
	// if Config.SoundcloudClientId == "" {
	// 	log.Fatalln("require soundcloud_client_id")
	// }

	fmt.Println("<Current Config>")
	fmt.Println("  * You can set your custom value by flag. see `$ gokuraku --help`")
	fmt.Printf("  gokuraku.Config.HttpPort: %s\n", Config.HttpPort)
	fmt.Printf("  gokuraku.Config.WebSocketPort: %s\n", Config.WebSocketPort)
	fmt.Printf("  gokuraku.Config.RedisAddr: %s\n", Config.RedisAddr)
	fmt.Printf("  gokuraku.Config.SoundcloudClientId: %s\n", Config.SoundcloudClientId)
	fmt.Println("</Current Config>")
	fmt.Println("")
}
