# Gokuraku

Gokuraku is music server.

![Logo](https://github.com/ToQoz/Gokuraku/raw/master/logo.png)

## Get Started

**if you don't have soundcloud, please register [here](http://soundcloud.com/you/apps/new)**

- Install. `go get github.com/stretchr/goweb`
- Run. `Gokuraku -soundcloud_client_id=YOUR_SOUNDCLOUD_CLIENT_ID`

## CommandLine options

```
$ go run main.go -help
  -p="9090": http server listen port
  -redis_addr=":6379": redis address(HOST:PORT)
  -redis_password="": redis password
  -soundcloud_client_id="": soundcloud client key
  -ws_p="9099": websocket server listen port
```
