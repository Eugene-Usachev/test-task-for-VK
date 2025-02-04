package handler

import (
	"strconv"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg"
	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg/model"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) initPingRoutes() {
	ping := h.app.Group("/ping")
	{
		ping.Get("/container/:id", h.GetPingsForContainer)
		ping.Post("/", h.StorePings)
	}
}

func (h *Handler) GetPingsForContainer(ctx fiber.Ctx) error {
	containerID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse container ID: " + err.Error())
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse `offset` query: " + err.Error())
	}

	pings, err := h.service.Ping.GetPingsForContainer(ctx.Context(), containerID, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to get pings: " + err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(pings)
}

func (h *Handler) StorePings(ctx fiber.Ctx) error {
	var pings []model.Ping
	if err := pkg.ParseJSON(ctx, &pings); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse request body: " + err.Error())
	}

	if err := h.service.Ping.StorePings(ctx.Context(), pings); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to store pings: " + err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}
