package service

import (
	"context"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/internal/repository"
	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg/model"
)

type Container interface {
	RegisterContainer(ctx context.Context, container *model.RegisterContainer) error
	GetContainers(ctx context.Context) ([]model.GetContainer, error)
	GetContainersWithLatestPing(
		ctx context.Context,
	) (successfulContainers []model.GetContainerWithLatestPing, invalidContainers []model.GetContainer, err error)
	UnregisterContainer(ctx context.Context, containerID int) error
}

type Ping interface {
	StorePings(ctx context.Context, pings []model.Ping) error
	GetPingsForContainer(ctx context.Context, containerID int, offset int) ([]model.Ping, error)
}

type Service struct {
	repository *repository.Repository

	Container
	Ping
}

func NewService() *Service {
	rep := repository.MustNewRepository()

	return &Service{
		repository: rep,

		Container: NewContainerService(rep),
		Ping:      NewPingService(rep),
	}
}
