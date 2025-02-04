package repository

import (
	"context"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/internal/config"
	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/internal/repository/postgres"
	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg/model"
)

type Container interface {
	RegisterContainer(ctx context.Context, container *model.RegisterContainer) error
	GetContainers(ctx context.Context) ([]model.GetContainer, error)
	GetContainersWithLatestPing(ctx context.Context) ([]model.GetContainerWithLatestPing, error)
	UnregisterContainer(ctx context.Context, containerID int) error
}

type Ping interface {
	StorePings(ctx context.Context, pings []model.Ping) error
	GetPingsForContainer(ctx context.Context, containerID int, offset int) ([]model.Ping, error)
}

type Repository struct {
	Container
	Ping
}

func MustNewRepository() *Repository {
	pool := postgres.MustNewPool(config.MustNewPostgresConfig())

	return &Repository{
		Container: NewContainerRepository(pool),
		Ping:      NewPingRepository(pool),
	}
}
