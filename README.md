# Gokuraku

![CI status](https://travis-ci.org/ToQoz/Gokuraku.png)

Gokuraku is music server.

![Screenshot](https://github.com/ToQoz/Gokuraku/raw/master/screenshot.png)

## Get Started

**if you don't have soundcloud, please register [here](http://soundcloud.com/you/apps/new)**

- Install. `go get github.com/ToQoz/Gokuraku`
- Run. `Gokuraku -soundcloud_client_id=YOUR_SOUNDCLOUD_CLIENT_ID`
- (Update) `go get -u github.com/ToQoz/Gokuraku`

## CommandLine options

```
$ Gokuraku -help
  -p="9090": http server listen port
  -redis_addr=":6379": redis address(HOST:PORT)
  -redis_password="": redis password
  -soundcloud_client_id="": soundcloud client key
  -ws_p="9099": websocket server listen port
```

## Requirements

- Go1.1
- Redis
- mercurial (require to install code.google.com/p/go.net/websocket)
- bzr (require to install github.com/stretchr/goweb)

---

![Logo](https://github.com/ToQoz/Gokuraku/raw/master/logo.png)
