package main

import "log"

// Hub maintains the set of active clients and broadcasts messages to the clients
type Hub struct {
	// Registered clients.
	clients map[*Client]bool
	// Inbound message from clients.
	broadcast chan []byte
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients.
	unregister chan *Client
	event      chan event
}

type eventType string
type event struct {
	clientID  string
	eventType eventType
}

func newHub() *Hub {
	return &Hub{
		clients:    map[*Client]bool{},
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		event:      make(chan event),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			log.Println("register", client)
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			log.Println("broadcast", string(message))
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
