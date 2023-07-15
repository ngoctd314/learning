package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

// Easier to get running with cors.
var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func main() {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	server.OnConnect("/", func(c socketio.Conn) error {
		log.Println("RUN")
		c.SetContext("")
		fmt.Println("connected:", c.ID())

		return nil
	})
	server.OnEvent("/", "connection", func(c socketio.Conn, msg string) {
		fmt.Println("new client connected", c.ID())
	})

	server.OnDisconnect("/", func(so socketio.Conn, reason string) {
		log.Println("Client disconnected:", so.ID())
	})

	server.OnError("/", func(c socketio.Conn, err error) {
		fmt.Println("meet error:", err)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	http.Handle("/socket.io/", server)
	// http.Handle("/", http.FileServer(http.Dir("./assets")))
	log.Println("Serving at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
