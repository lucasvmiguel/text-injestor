package api

import "net/http"

// Config struct to config package
type Config struct {
	Port     string
	Handlers map[string]func(w http.ResponseWriter, r *http.Request)
}
