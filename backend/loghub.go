package main

import "sync"

type logHub struct {
	hubsLck sync.RWMutex
	hubs    map[string]*Hub
}

func (lhub *logHub) add(key string, hub *Hub) {
	lhub.hubsLck.Lock()
	lhub.hubs[key] = hub
	go lhub.hubs[key].run()
	lhub.hubsLck.Unlock()
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*LogStreamer]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *LogStreamer

	// Unregister requests from clients.
	unregister chan *LogStreamer
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *LogStreamer),
		unregister: make(chan *LogStreamer),
		clients:    make(map[*LogStreamer]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
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
