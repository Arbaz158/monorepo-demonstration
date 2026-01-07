package config

import "os"

// Config holds minimal configuration for a service.
type Config struct {
	ServiceName string
	Port        string
}

// Load builds a Config using environment variables and sensible defaults.
func Load(service string) Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{ServiceName: service, Port: port}
}
