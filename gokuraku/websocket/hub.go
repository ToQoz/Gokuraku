package websocket

import (
	"code.google.com/p/go.net/websocket"
	"github.com/ToQoz/Gokuraku/gokuraku"
	"github.com/ToQoz/Gokuraku/gokuraku/models/track"
	"log"
	"net/http"
)

type Hub struct {
	Conns            []*Conn
	AddConn          chan *Conn
	RemoveConn       chan *Conn
	AddTrack         chan *Conn
	UpdatedConnState chan *Conn
}

func Run() {
	hub := NewHub()
	hub.Run()
}

func NewHub() Hub {
	return Hub{
		Conns:            make([]*Conn, 0),
		AddConn:          make(chan *Conn),
		RemoveConn:       make(chan *Conn),
		AddTrack:         make(chan *Conn),
		UpdatedConnState: make(chan *Conn),
	}
}

func (s *Hub) Run() {
	http.Handle("/ws", s.connHandler())

	go func() {
		log.Printf("Gokuraku WebSocket Server: 0.0.0.0:%s", gokuraku.Config.WebSocketPort)
		err := http.ListenAndServe("0.0.0.0:"+gokuraku.Config.WebSocketPort, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	for {
		select {
		case <-s.UpdatedConnState:
			if s.IsNoonePlaying() == true {
				s.broadcastTrack()
			}
		case conn := <-s.AddConn:
			s.addConn(conn)
			log.Println("Number of clients connected ...", len(s.Conns))
		case conn := <-s.RemoveConn:
			s.removeConn(conn)
			log.Println("Number of clients connected ...", len(s.Conns))
		case <-s.AddTrack:
			if s.IsNoonePlaying() == true {
				s.broadcastTrack()
			}
		}
	}
}

func (s *Hub) addConn(c *Conn) {
	s.Conns = append(s.Conns, c)

	if s.IsNoonePlaying() == true {
		s.broadcastTrack()
	}
}

func (s *Hub) removeConn(c *Conn) {
	for i, t := range s.Conns {
		if t == c {
			s.Conns = append(s.Conns[:i], s.Conns[i+1:]...)
			break
		}
	}

	if s.IsNoonePlaying() == true {
		s.broadcastTrack()
	}
}

func (s *Hub) IsNoonePlaying() bool {
	if len(s.Conns) == 0 {
		return false
	}

	for _, c := range s.Conns {
		if c.IsPlaying() {
			return false
		}
	}

	return true
}

func (s *Hub) broadcastTrack() {
	currentTrack, err := track.Next()

	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, conn := range s.Conns {
		if conn.IsReadyToPlay() {
			conn.Play <- currentTrack
		}
	}
}

func (s *Hub) connHandler() websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		var err error

		defer func() {
			err = ws.Close()
			if err != nil {
				log.Println("Websocket could not be closed", err.Error())
			}
		}()

		conn := NewConn(ws, s)

		go conn.WritePump()
		currentTrack, err := track.GetCurrent()
		if err == nil {
			conn.Play <- currentTrack
		} else {
			log.Println(err.Error())
		}
		s.AddConn <- conn
		conn.ReadPump()
		defer ws.Close()
	})
}
