package main

import (
	"log"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/internal/config"
	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/internal/handler"
	servicepkg "github.com/Eugene-Usachev/test-task-for-VK/backend/src/internal/service"
)

func main() {
	cfg := config.MustNewAppConfig()

	// service is not created in handler.NewHTTPHandler, because handler.Handler supports HTTP only,
	// if we will create, for example, gRPC handler it will use the same service.
	service := servicepkg.NewService()

	if err := handler.NewHTTPHandler(service).Run(cfg.Addr()); err != nil {
		log.Panicf("Failed to start server: %v", err)
	}
}
