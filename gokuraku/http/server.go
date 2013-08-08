package http

import (
	"github.com/ToQoz/Gokuraku/gokuraku"
	"github.com/stretchr/goweb"
	"log"
	"net"
	httplib "net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	// HTTP
	s := &httplib.Server{
		Addr:           gokuraku.Config.HttpAddr,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, listenErr := net.Listen("tcp", gokuraku.Config.HttpAddr)

	if listenErr != nil {
		log.Fatalf("Could not listen: %s", gokuraku.Config.HttpAddr)
	}

	go func() {
		for _ = range c {
			// sig is a ^C, handle it
			log.Print("Stopping the server...")
			listener.Close()

			log.Print("Tearing down...")
			log.Fatal("Finished - bye bye.  ;-)")

		}
	}()

	log.Printf("Gokuraku HTTP Server: %s", gokuraku.Config.HttpAddr)
	log.Fatalf("Error in Serve: %s", s.Serve(listener))
}
