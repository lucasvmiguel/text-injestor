package api

import (
	"net/http"

	"github.com/pkg/errors"
)

// Client struct to use package
type Client struct {
	config Config
}

// New receives the configuration
func New(config Config) (Client, error) {
	if config.Port == "" {
		return Client{}, errors.New("invalid port")
	}

	if config.Handlers == nil || len(config.Handlers) < 1 {
		return Client{}, errors.New("empty handlers map")
	}

	for route, handler := range config.Handlers {
		http.HandleFunc(route, handler)
	}

	return Client{config: config}, nil
}

// Run http server
func (c *Client) Run() {
	http.ListenAndServe(c.config.Port, nil)
}
