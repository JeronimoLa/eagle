package client

import "net/http"

type Config struct {
	HTTPClient *http.Client
}

func New(cfg *Config) *Config {
	return &Config{
		HTTPClient: cfg.HTTPClient,
		// Endpoint: baseURL,
	}
}
