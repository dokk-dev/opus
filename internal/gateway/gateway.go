package gateway

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/dokk-dev/opus/internal/config"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: Implement proper origin checking in production
	},
}

// Message represents a WebSocket message
type Message struct {
	Type      string          `json:"type"`
	Channel   string          `json:"channel,omitempty"`
	From      string          `json:"from,omitempty"`
	Content   string          `json:"content,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
	Timestamp int64           `json:"timestamp,omitempty"`
}

// Client represents a connected WebSocket client
type Client struct {
	ID       string
	Conn     *websocket.Conn
	Gateway  *Gateway
	Send     chan []byte
	Channels map[string]bool
	Role     string // manager, assistant_manager, admin
}

// Gateway manages WebSocket connections and message routing
type Gateway struct {
	config     *config.Config
	clients    map[string]*Client
	channels   map[string]map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
	mu         sync.RWMutex
}

// New creates a new Gateway instance
func New(cfg *config.Config) *Gateway {
	gw := &Gateway{
		config:     cfg,
		clients:    make(map[string]*Client),
		channels:   make(map[string]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}

	go gw.run()
	return gw
}

func (gw *Gateway) run() {
	for {
		select {
		case client := <-gw.register:
			gw.mu.Lock()
			gw.clients[client.ID] = client
			gw.mu.Unlock()
			log.Printf("Client connected: %s", client.ID)

		case client := <-gw.unregister:
			gw.mu.Lock()
			if _, ok := gw.clients[client.ID]; ok {
				delete(gw.clients, client.ID)
				close(client.Send)
				// Remove from all channels
				for channel := range client.Channels {
					if clients, ok := gw.channels[channel]; ok {
						delete(clients, client)
					}
				}
			}
			gw.mu.Unlock()
			log.Printf("Client disconnected: %s", client.ID)

		case msg := <-gw.broadcast:
			gw.broadcastToChannel(msg)
		}
	}
}

func (gw *Gateway) broadcastToChannel(msg *Message) {
	gw.mu.RLock()
	defer gw.mu.RUnlock()

	if msg.Channel == "" {
		// Broadcast to all clients
		for _, client := range gw.clients {
			gw.sendToClient(client, msg)
		}
	} else {
		// Broadcast to specific channel
		if clients, ok := gw.channels[msg.Channel]; ok {
			for client := range clients {
				gw.sendToClient(client, msg)
			}
		}
	}
}

func (gw *Gateway) sendToClient(client *Client, msg *Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return
	}

	select {
	case client.Send <- data:
	default:
		// Client buffer full, skip
		log.Printf("Client %s buffer full, skipping message", client.ID)
	}
}

// HandleWebSocket handles WebSocket upgrade and connection
func (gw *Gateway) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := &Client{
		ID:       generateClientID(),
		Conn:     conn,
		Gateway:  gw,
		Send:     make(chan []byte, 256),
		Channels: make(map[string]bool),
	}

	gw.register <- client

	go client.writePump()
	go client.readPump()
}

// Broadcast sends a message to all clients or a specific channel
func (gw *Gateway) Broadcast(msg *Message) {
	gw.broadcast <- msg
}

// JoinChannel adds a client to a channel
func (gw *Gateway) JoinChannel(client *Client, channel string) {
	gw.mu.Lock()
	defer gw.mu.Unlock()

	if _, ok := gw.channels[channel]; !ok {
		gw.channels[channel] = make(map[*Client]bool)
	}
	gw.channels[channel][client] = true
	client.Channels[channel] = true
}

func (c *Client) readPump() {
	defer func() {
		c.Gateway.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		// Handle message based on type
		c.handleMessage(&msg)
	}
}

func (c *Client) writePump() {
	defer c.Conn.Close()

	for message := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Write error: %v", err)
			return
		}
	}
}

func (c *Client) handleMessage(msg *Message) {
	switch msg.Type {
	case "join":
		c.Gateway.JoinChannel(c, msg.Channel)
	case "chat":
		// Route to AI handler
		msg.From = c.ID
		c.Gateway.Broadcast(msg)
	case "ping":
		response := &Message{Type: "pong"}
		data, _ := json.Marshal(response)
		c.Send <- data
	}
}

func generateClientID() string {
	// Simple ID generation - use UUID in production
	return "client_" + randomString(8)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[i%len(letters)]
	}
	return string(b)
}
