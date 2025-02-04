package client

import (
	"context"

	"github.com/Eugene-Usachev/test-task-for-VK/pinger/src/pkg/model"
)

type Client interface {
	GetContainers(ctx context.Context) ([]model.GetContainer, error)
	StorePings(ctx context.Context, pings []model.Ping) error
}
