package service

import (
	"context"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/internal/repository"
	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg/model"
)

type PingService struct {
	repository repository.Ping
}

var _ Ping = (*PingService)(nil)

func NewPingService(repository repository.Ping) *PingService {
	return &PingService{
		repository: repository,
	}
}

func (p *PingService) StorePings(ctx context.Context, pings []model.Ping) error {
	return p.repository.StorePings(ctx, pings)
}

func (p *PingService) GetPingsForContainer(
	ctx context.Context,
	containerID int,
	offset int,
) ([]model.Ping, error) {
	return p.repository.GetPingsForContainer(ctx, containerID, offset)
}
