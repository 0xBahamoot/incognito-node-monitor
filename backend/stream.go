package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 5 * time.Second
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// LogStreamer is a middleman between the websocket connection and the hub.
type LogStreamer struct {
	hub *Hub

	id Hash
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
func (c *LogStreamer) readPump() {
	defer func() {
		if c.hub != nil {
			c.hub.unregister <- c
		}
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
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (c *LogStreamer) writePump() {
	defer func() {
		if c.hub != nil {
			c.hub.unregister <- c
		}
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(err)
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

// streamlogWs handles websocket requests from the peer.
func streamlogWs(hub *Hub, w http.ResponseWriter, r *http.Request, preStreamLogs []string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	sendBuffer := 256
	client := &LogStreamer{hub: hub, conn: conn, send: make(chan []byte, sendBuffer), id: HashH([]byte(r.RemoteAddr))}
	go client.writePump()

	if len(preStreamLogs) > 0 {
		for i := len(preStreamLogs) - 1; i >= 0; i-- {
			client.send <- []byte(preStreamLogs[i])
		}
	}

	client.hub.register <- client
	// go client.readPump()
}

func streamOnceWs(w http.ResponseWriter, r *http.Request, streamLogs []string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	sendBuffer := 256
	client := &LogStreamer{hub: nil, conn: conn, send: make(chan []byte, sendBuffer), id: HashH([]byte(r.RemoteAddr))}
	go client.writePump()

	if len(streamLogs) > 0 {
		for i := 0; i <= len(streamLogs)-1; i++ {
			client.send <- []byte(streamLogs[i])
		}
	}
	close(client.send)
}
