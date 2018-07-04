package main

import (
	"net/http"

	"github.com/lucasvmiguel/text-injestor/api"
	"github.com/lucasvmiguel/text-injestor/handlers"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type appConfig struct {
	http httpConfig
}

type httpConfig struct {
	port     string
	handlers httpHandlersConfig
}

type httpHandlersConfig struct {
	stats string
}

func main() {
	// load config from a config file (config.toml)
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	// mount the api config
	apiConfig := api.Config{
		Port: config.http.port,
		Handlers: map[string]func(w http.ResponseWriter, r *http.Request){
			config.http.handlers.stats: handlers.Stats,
		},
	}

	// create a new api instance with the api config
	apiClient, err := api.New(apiConfig)
	if err != nil {
		panic(errors.Wrap(err, "error to initialize api"))
	}

	// run the api
	err = apiClient.Run()
	if err != nil {
		panic(errors.Wrap(err, "error to run the api"))
	}
}

func loadConfig() (appConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		return appConfig{}, errors.Wrap(err, "error to read config")
	}

	return appConfig{
		httpConfig{
			port: viper.GetString("http.port"),
			handlers: httpHandlersConfig{
				stats: viper.GetString("http.handlers.stats"),
			},
		},
	}, nil
}
