package pinger

import "github.com/Eugene-Usachev/test-task-for-VK/pinger/src/pkg/model"

type Pinger interface {
	PingEachContainer(containers []model.GetContainer) []model.Ping
}
