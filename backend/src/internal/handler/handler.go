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

	handler.initRoutes()

	return handler
}

func (h *Handler) initRoutes() {
	h.initContainerRoutes()
	h.initPingRoutes()
}

func (h *Handler) Run(addr string) error {
	return h.app.Listen(addr)
}
