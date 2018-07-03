package main

import (
	"net/http"

	"github.com/lucasvmiguel/text-injestor/api"
	"github.com/lucasvmiguel/text-injestor/handlers"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type config struct {
	http httpConfig
}

type httpConfig struct {
	port     string
	handlers httpHandlersConfig
}

type httpHandlersConfig struct {
	stats string
}

func loadConfig() (config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return config{}, errors.Wrap(err, "error to read config")
	}

	return config{
		httpConfig{
			port: viper.GetString("http.port"),
			handlers: httpHandlersConfig{
				stats: viper.GetString("http.handlers.stats"),
			},
		},
	}, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	apiConfig := api.Config{
		Port: config.http.port,
		Handlers: map[string]func(w http.ResponseWriter, r *http.Request){
			config.http.handlers.stats: handlers.Stats,
		},
	}

	apiClient, err := api.New(apiConfig)
	if err != nil {
		panic(errors.Wrap(err, "error to initialize api"))
	}

	apiClient.Run()
}
