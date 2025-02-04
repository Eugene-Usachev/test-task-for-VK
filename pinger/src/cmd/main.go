package main

import (
	"context"
	"log"
	"time"

	clientpkg "github.com/Eugene-Usachev/test-task-for-VK/pinger/src/internal/client"
	"github.com/Eugene-Usachev/test-task-for-VK/pinger/src/internal/config"
	pingerpkg "github.com/Eugene-Usachev/test-task-for-VK/pinger/src/internal/pinger"
)

func tryPing(pinger pingerpkg.Pinger, client clientpkg.Client) {
	// It is not difficult to add retry logic here,
	// but it looks like if pinger cannot connect to the server
	// then it is better to wait whole interval because the server is down and pinger
	// don't need to provide a minimal latency.
	ctx := context.Background()

	containers, err := client.GetContainers(ctx)
	if err != nil {
		log.Printf("Failed to get containers: %v", err)

		return
	}

	pinger.PingEachContainer(containers)

	err = client.StorePings(ctx, pinger.PingEachContainer(containers))
	if err != nil {
		log.Printf("Failed to store pings: %v", err)

		return
	}

	log.Println("Successfully stored pings, date: ", time.Now().Format(time.RFC3339))
}

func main() {
	cfg := config.MustNewAppConfig()

	pinger := pingerpkg.NewGoPinger(cfg.GetTimeout(), cfg.GetTries())
	client := clientpkg.NewHTTPClient(cfg.BackendAddr())

	for {
		tryPing(pinger, client)

		time.Sleep(cfg.GetInterval())
	}
}
