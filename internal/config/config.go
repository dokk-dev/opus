package config

import (
	"os"
)

type Config struct {
	// Server settings
	ServerAddr string

	// AI settings
	OllamaURL       string
	OllamaModel     string
	ClaudeAPIKey    string
	ClaudeFallback  bool

	// Database
	DatabaseURL string

	// Security
	JWTSecret   string
	CORSOrigins []string
}

func Load() (*Config, error) {
	cfg := &Config{
		ServerAddr:     getEnv("SERVER_ADDR", ":8080"),
		OllamaURL:      getEnv("OLLAMA_URL", "http://localhost:11434"),
		OllamaModel:    getEnv("OLLAMA_MODEL", "llama3"),
		ClaudeAPIKey:   getEnv("CLAUDE_API_KEY", ""),
		ClaudeFallback: getEnv("CLAUDE_FALLBACK", "true") == "true",
		DatabaseURL:    getEnv("DATABASE_URL", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		CORSOrigins:    []string{getEnv("CORS_ORIGINS", "http://localhost:5173")},
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
