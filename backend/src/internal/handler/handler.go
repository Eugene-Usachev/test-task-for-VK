package handler

import (
	servicepkg "github.com/Eugene-Usachev/test-task-for-VK/backend/src/internal/service"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	app     *fiber.App
	service *servicepkg.Service
}

func NewHTTPHandler(service *servicepkg.Service) *Handler {
	handler := &Handler{
		app:     fiber.New(),
		service: service,
	}

	handler.initCORSMiddleware()
	handler.initRoutes()

	return handler
}

func (h *Handler) initCORSMiddleware() {
	h.app.Use(func(c fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}

		return c.Next()
	})
}

func (h *Handler) initRoutes() {
	h.initContainerRoutes()
	h.initPingRoutes()
}

func (h *Handler) Run(addr string) error {
	return h.app.Listen(addr)
}
