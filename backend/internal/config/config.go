package config

import "os"

type Config struct {
	Port        string
	FrontendURL string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	return &Config{
		Port:        port,
		FrontendURL: frontendURL,
	}
}
