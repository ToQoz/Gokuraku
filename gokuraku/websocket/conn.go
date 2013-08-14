package websocket

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"github.com/ToQoz/Gokuraku/gokuraku/models/track"
	"log"
)

const (
	WaitingState = "waiting"
	PlayingState = "playing"
)

type Conn struct {
	disconnect chan bool
	ws         *websocket.Conn
	hub        *Hub
	Play       chan *track.CurrentTrack
	State      string
}

func NewConn(ws *websocket.Conn, hub *Hub) *Conn {
	return &Conn{
		ws:         ws,
		hub:        hub,
		disconnect: make(chan bool),
		Play:       make(chan *track.CurrentTrack),
		State:      WaitingState,
	}
}

func (c *Conn) IsReadyToPlay() bool {
	return c.State == WaitingState
}

func (c *Conn) IsPlaying() bool {
	return c.State == PlayingState
}

func (c *Conn) updateState(state string) {
	log.Printf("Update Conn#State: %s -> %s", c.State, state)
	c.State = state
	c.hub.UpdatedConnState <- c
}

func (c *Conn) WritePump() {
	for {
		select {
		case <-c.disconnect:
			c.hub.RemoveConn <- c
			c.disconnect <- true
			return
		case track := <-c.Play:
			playMsg := newPlayMessage(track)
			playMsgJson, err := json.Marshal(playMsg)

			c.updateState(PlayingState)
			websocket.Message.Send(c.ws, string(playMsgJson))

			if err != nil {
				log.Println(err.Error())
				c.hub.RemoveConn <- c
				c.disconnect <- true
				return
			}
		}
	}
}

func (c *Conn) ReadPump() {
	for {
		select {
		case <-c.disconnect:
			c.hub.RemoveConn <- c
			c.disconnect <- true
			return
		default:
			var err error
			var msg []byte
			err = websocket.Message.Receive(c.ws, &msg)
			var msgMap map[string]interface{}
			json.Unmarshal(msg, &msgMap)
			t := msgMap["Type"]

			if err != nil {
				log.Println(err.Error())
				c.hub.RemoveConn <- c
				c.disconnect <- true
				return
			} else {
				switch {
				case t == WaitingState:
					c.updateState(WaitingState)
				case t == PlayingState:
					c.updateState(PlayingState)
				}
			}
		}
	}
}

type playMessage struct {
	Type  string
	Track *track.CurrentTrack
}

func newPlayMessage(track *track.CurrentTrack) *playMessage {
	return &playMessage{Type: "play", Track: track}
}
