package config

import (
	"log"
	"os"
)

type PostgresConfig struct {
	user     string
	password string
	host     string
	port     string
	database string
}

func MustNewPostgresConfig() PostgresConfig {
	var config PostgresConfig

	isValid := true

	config.user = os.Getenv("POSTGRES_USER")
	if config.user == "" {
		isValid = false

		log.Println("POSTGRES_USER is not set")
	}

	config.password = os.Getenv("POSTGRES_PASSWORD")
	if config.password == "" {
		isValid = false

		log.Println("POSTGRES_PASSWORD is not set")
	}

	config.host = os.Getenv("POSTGRES_HOST")
	if config.host == "" {
		isValid = false

		log.Println("POSTGRES_HOST is not set")
	}

	config.port = os.Getenv("POSTGRES_PORT")
	if config.port == "" {
		isValid = false

		log.Println("POSTGRES_PORT is not set")
	}

	config.database = os.Getenv("POSTGRES_DATABASE")
	if config.database == "" {
		isValid = false

		log.Println("POSTGRES_DATABASE is not set")
	}

	if !isValid {
		log.Fatal("Invalid config, read logs above")
	}

	return config
}

func (config PostgresConfig) User() string {
	return config.user
}

func (config PostgresConfig) Password() string {
	return config.password
}

func (config PostgresConfig) Host() string {
	return config.host
}

func (config PostgresConfig) Port() string {
	return config.port
}

func (config PostgresConfig) Database() string {
	return config.database
}
