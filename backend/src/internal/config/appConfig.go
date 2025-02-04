package config

import (
	"log"
	"os"
)

type AppConfig struct {
	addr string
}

func MustNewAppConfig() AppConfig {
	var config AppConfig

	isValid := true

	config.addr = os.Getenv("ADDR")
	if config.addr == "" {
		isValid = false

		log.Println("ADDR is not set")
	}

	if !isValid {
		log.Fatal("Invalid config, read logs above")
	}

	return config
}

func (config AppConfig) Addr() string {
	return config.addr
}
