package service

import (
	"context"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/internal/repository"
	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg/model"
)

type ContainerService struct {
	repository repository.Container
}

var _ Container = (*ContainerService)(nil)

func NewContainerService(repository repository.Container) *ContainerService {
	return &ContainerService{
		repository: repository,
	}
}

func (c ContainerService) RegisterContainer(ctx context.Context, container *model.RegisterContainer) error {
	return c.repository.RegisterContainer(ctx, container)
}

func (c ContainerService) GetContainers(ctx context.Context) ([]model.GetContainer, error) {
	return c.repository.GetContainers(ctx)
}

func (c ContainerService) GetContainersWithLatestPing(
	ctx context.Context,
) ([]model.GetContainerWithLatestPing, error) {
	return c.repository.GetContainersWithLatestPing(ctx)
}

func (c ContainerService) UnregisterContainer(ctx context.Context, containerID int) error {
	return c.repository.UnregisterContainer(ctx, containerID)
}
