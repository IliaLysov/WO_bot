package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"tg_weather/pkg/ow"
	"tg_weather/pkg/tg"
)

type Config struct {
	Tg tg.Config
	OW ow.Config
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		return config, fmt.Errorf("Error loading .env file: %w", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("Error loading .env file: %w", err)
	}

	return config, nil
}
