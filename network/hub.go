package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	// ClientID
	ID  string
	hub *Hub
	// The websocket connection
	conn *websocket.Conn
	// Buffered channel of outbound messages
	send chan []byte
}

// readPump pumps message from the ws connection to the hub
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		log.Println("recv message: ", string(message))
		// c.hub.event <- &Event{
		// 	ClientID: c.ID,
		// 	Message:  EventFromByte(message),
		// }
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			fmt.Println("receive message", string(message))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}

}

type EventMessage string

const (
	QRCodeScanned EventMessage = "qrcode-scanned"
	UnKnownEvent  EventMessage = "unknown"
)

func ValidateEvent(p []byte) bool {
	return true
	// msg := EventMessage(p)
	// fmt.Println("msg", msg)
	// switch msg {
	// case QRCodeScanned:
	// 	return true
	// default:
	// 	return false
	// }
}

type Event struct {
	ClientID string
	Message  []byte
}

// Hub maintains the set of active clients
type Hub struct {
	// Registered clients.
	clients map[string]*Client
	// Register request from client
	register chan *Client
	// Unrester request from client
	unregister chan *Client
	event      chan *Event
}

func NewHub() *Hub {
	return &Hub{
		clients:    map[string]*Client{},
		register:   make(chan *Client),
		unregister: make(chan *Client),
		event:      make(chan *Event),
	}
}

func (h *Hub) run() {
	// blocking mode
	for {
		select {
		case client := <-h.register:
			log.Println("register new client", client.ID)
			h.clients[client.ID] = client
		case client := <-h.unregister:
			log.Println("un register new client", client.ID)
			if _, registed := h.clients[client.ID]; registed {
				// delete client
				delete(h.clients, client.ID)
				// close client
			}
		case evt := <-h.event:
			log.Println("receive client event", evt.ClientID)
			// get client for event
			if client, registed := h.clients[evt.ClientID]; registed {
				// valide event
				if ValidateEvent(evt.Message) {
					client.send <- evt.Message
				}
			}
		}
	}
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	qrcodeID := r.URL.Query().Get("qrcode_id")
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{ID: qrcodeID, hub: hub, conn: conn, send: make(chan []byte, 255)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
