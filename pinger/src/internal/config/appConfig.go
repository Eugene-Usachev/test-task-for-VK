package config

import (
	"github.com/goccy/go-json"
	"log"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	backendAddr    string
	timeout        time.Duration
	tries          int
	interval       time.Duration
	containerAddrs []string
}

func MustNewAppConfig() AppConfig {
	var (
		config AppConfig
		err    error
	)

	isValid := true

	config.backendAddr = os.Getenv("BACKEND_ADDR")
	if config.backendAddr == "" {
		isValid = false

		log.Println("BACKEND_ADDR is not set")
	}

	timeoutStr := os.Getenv("TIMEOUT_MS")
	if timeoutStr == "" {
		isValid = false

		log.Println("TIMEOUT_MS is not set")
	} else {
		config.timeout, err = time.ParseDuration(timeoutStr + "ms")
		if err != nil {
			isValid = false

			log.Printf("Failed to parse TIMEOUT_MS: %v", err)
		}
	}

	triesStr := os.Getenv("TRIES")
	if triesStr == "" {
		isValid = false

		log.Println("TRIES is not set")
	} else {
		config.tries, err = strconv.Atoi(triesStr)
		if err != nil {
			isValid = false

			log.Printf("Failed to parse TRIES: %v", err)
		}
	}

	intervalStr := os.Getenv("INTERVAL_MS")
	if intervalStr == "" {
		isValid = false

		log.Println("INTERVAL_MS is not set")
	} else {
		config.interval, err = time.ParseDuration(intervalStr + "ms")
		if err != nil {
			isValid = false

			log.Printf("Failed to parse INTERVAL_MS: %v", err)
		}
	}

	containerAddrs := os.Getenv("CONTAINER_ADDRS")
	if containerAddrs == "" {
		isValid = false
	} else {
		if err = json.Unmarshal([]byte(containerAddrs), &config.containerAddrs); err != nil {
			isValid = false

			log.Printf("Failed to parse CONTAINER_ADDRS: %v", err)
		}
	}

	if !isValid {
		log.Fatal("Invalid config, read logs above")
	}

	return config
}

func (config *AppConfig) BackendAddr() string {
	return config.backendAddr
}

func (config *AppConfig) GetTimeout() time.Duration {
	return config.timeout
}

func (config *AppConfig) GetTries() int {
	return config.tries
}

func (config *AppConfig) GetInterval() time.Duration {
	return config.interval
}

func (config *AppConfig) GetContainerAddrs() []string {
	return config.containerAddrs
}
