package handler

import (
	"strconv"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg"
	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg/model"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) initContainerRoutes() {
	container := h.app.Group("/container")
	{
		container.Get("/id_and_ip_address_only", h.GetContainerIDsAndIPAddressesOnly)
		container.Get("/with_latest_ping", h.GetContainersWithLatestPing)
		container.Post("/one", h.RegisterContainer)
		container.Post("/many", h.RegisterManyContainers)
		container.Delete("/:id", h.UnregisterContainer)
	}
}

func (h *Handler) GetContainerIDsAndIPAddressesOnly(ctx fiber.Ctx) error {
	containers, err := h.service.Container.GetContainers(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to get containers: " + err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(containers)
}

func (h *Handler) GetContainersWithLatestPing(ctx fiber.Ctx) error {
	successfulContainers, invalidContainers, err := h.service.Container.GetContainersWithLatestPing(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to get containers: " + err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(struct {
		SuccessfulContainers []model.GetContainerWithLatestPing `json:"successfulContainers"`
		InvalidContainers    []model.GetContainer               `json:"invalidContainers"`
	}{
		SuccessfulContainers: successfulContainers,
		InvalidContainers:    invalidContainers,
	})
}

func (h *Handler) RegisterContainer(ctx fiber.Ctx) error {
	var container model.RegisterContainer
	if err := pkg.ParseJSON(ctx, &container); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse request body: " + err.Error())
	}

	if err := h.service.Container.RegisterContainer(ctx.Context(), &container); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to register container: " + err.Error())
	}

	return ctx.Status(fiber.StatusOK).SendString("Container registered successfully")
}

func (h *Handler) RegisterManyContainers(ctx fiber.Ctx) error {
	var containers []model.RegisterContainer
	if err := pkg.ParseJSON(ctx, &containers); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse request body: " + err.Error())
	}

	for i := range containers {
		if err := h.service.Container.RegisterContainer(ctx.Context(), &containers[i]); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to register container: " + err.Error())
		}
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h *Handler) UnregisterContainer(ctx fiber.Ctx) error {
	containerID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse container ID: " + err.Error())
	}

	if err = h.service.Container.UnregisterContainer(ctx.Context(), containerID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to unregister container: " + err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}
