package client

import "net/http"

type Config struct {
	HTTPClient *http.Client
	UserAgent  string
}

func New(cfg *Config) *Config {
	return &Config{
		HTTPClient: cfg.HTTPClient,
		UserAgent:  cfg.UserAgent,
	}
}


