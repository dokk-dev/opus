package api

import (
	"encoding/json"
	"net/http"

	"github.com/dokk-dev/opus/internal/ai"
	"github.com/dokk-dev/opus/internal/config"
	"github.com/dokk-dev/opus/internal/gateway"
)

type Router struct {
	mux      *http.ServeMux
	config   *config.Config
	gateway  *gateway.Gateway
	aiRouter *ai.Router
}

func NewRouter(cfg *config.Config, gw *gateway.Gateway, aiRouter *ai.Router) *Router {
	r := &Router{
		mux:      http.NewServeMux(),
		config:   cfg,
		gateway:  gw,
		aiRouter: aiRouter,
	}

	r.setupRoutes()
	return r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// CORS middleware
	w.Header().Set("Access-Control-Allow-Origin", r.config.CORSOrigins[0])
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	r.mux.ServeHTTP(w, req)
}

func (r *Router) setupRoutes() {
	// Health check
	r.mux.HandleFunc("GET /health", r.healthCheck)

	// WebSocket endpoint
	r.mux.HandleFunc("GET /ws", r.gateway.HandleWebSocket)

	// API v1
	r.mux.HandleFunc("GET /api/v1/status", r.getStatus)
	r.mux.HandleFunc("POST /api/v1/chat", r.handleChat)

	// Department endpoints
	r.mux.HandleFunc("GET /api/v1/departments", r.getDepartments)
	r.mux.HandleFunc("GET /api/v1/departments/{dept}/inventory", r.getDepartmentInventory)
	r.mux.HandleFunc("GET /api/v1/departments/{dept}/schedule", r.getDepartmentSchedule)

	// Alerts
	r.mux.HandleFunc("GET /api/v1/alerts", r.getAlerts)
}

func (r *Router) healthCheck(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func (r *Router) getStatus(w http.ResponseWriter, req *http.Request) {
	status := map[string]interface{}{
		"server":  "running",
		"ai":      "ollama",
		"model":   r.config.OllamaModel,
		"version": "0.1.0",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

// ChatRequest represents an incoming chat message
type ChatRequest struct {
	Message string       `json:"message"`
	History []ai.Message `json:"history,omitempty"`
}

// ChatResponse represents the AI response
type ChatResponse struct {
	Response   string `json:"response"`
	Department string `json:"department"`
	Model      string `json:"model"`
}

func (r *Router) handleChat(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var chatReq ChatRequest
	if err := json.NewDecoder(req.Body).Decode(&chatReq); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if chatReq.Message == "" {
		http.Error(w, `{"error": "Message is required"}`, http.StatusBadRequest)
		return
	}

	// Process through AI router
	response, dept, err := r.aiRouter.ProcessQuery(req.Context(), chatReq.Message, chatReq.History)
	if err != nil {
		// Return error but don't expose internal details
		http.Error(w, `{"error": "Failed to process message. Is Ollama running?"}`, http.StatusServiceUnavailable)
		return
	}

	json.NewEncoder(w).Encode(ChatResponse{
		Response:   response,
		Department: string(dept),
		Model:      r.config.OllamaModel,
	})
}

func (r *Router) getDepartments(w http.ResponseWriter, req *http.Request) {
	// Demo data
	departments := []map[string]interface{}{
		{"id": "dairy", "name": "Dairy", "manager": "Demo Manager"},
		{"id": "produce", "name": "Produce", "manager": "Demo Manager"},
		{"id": "meat", "name": "Meat", "manager": "Demo Manager"},
		{"id": "bakery", "name": "Bakery", "manager": "Demo Manager"},
		{"id": "deli", "name": "Deli", "manager": "Demo Manager"},
		{"id": "grocery", "name": "Grocery", "manager": "Demo Manager"},
		{"id": "frontend", "name": "Front End", "manager": "Demo Manager"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(departments)
}

func (r *Router) getDepartmentInventory(w http.ResponseWriter, req *http.Request) {
	dept := req.PathValue("dept")
	// Demo data
	inventory := map[string]interface{}{
		"department": dept,
		"items":      []map[string]interface{}{},
		"lastUpdated": "2024-01-01T00:00:00Z",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}

func (r *Router) getDepartmentSchedule(w http.ResponseWriter, req *http.Request) {
	dept := req.PathValue("dept")
	// Demo data
	schedule := map[string]interface{}{
		"department": dept,
		"shifts":     []map[string]interface{}{},
		"week":       "current",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedule)
}

func (r *Router) getAlerts(w http.ResponseWriter, req *http.Request) {
	// Demo data
	alerts := []map[string]interface{}{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alerts)
}
