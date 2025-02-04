package pinger

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Eugene-Usachev/test-task-for-VK/pinger/src/pkg/model"
	probing "github.com/prometheus-community/pro-bing"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GoPinger struct {
	wg      *sync.WaitGroup
	slice   []model.Ping
	tail    atomic.Int64
	timeout time.Duration
	tries   int
}

var _ Pinger = (*GoPinger)(nil)

func NewGoPinger(timeout time.Duration, tries int) *GoPinger {
	return &GoPinger{
		wg:      &sync.WaitGroup{},
		slice:   nil,
		tail:    atomic.Int64{},
		timeout: timeout,
		tries:   tries,
	}
}

func (p *GoPinger) PingEachContainer(containers []model.GetContainer) []model.Ping {
	p.wg.Add(len(containers))
	p.tail.Store(0)

	if len(p.slice) < len(containers) {
		p.slice = append(p.slice, make([]model.Ping, len(containers)-len(p.slice))...)
	}

	for i := range containers {
		go p.ping(&containers[i])
	}

	p.wg.Wait()

	return p.slice
}

// It also calls wg.Done().
func (p *GoPinger) ping(container *model.GetContainer) {
	defer p.wg.Done()

	pinger, err := probing.NewPinger(container.GetIpAddress())
	if err != nil {
		p.writeResult(container.GetId(), 0, false)
		log.Printf(
			"Failed to create pinger for container with id `%d` and ip address `%s`: %v",
			container.GetId(), container.GetIpAddress(), err,
		)

		return
	}

	pinger.Timeout = p.timeout
	pinger.Count = p.tries

	err = pinger.Run()
	if err != nil {
		p.writeResult(container.GetId(), 0, false)
		log.Printf(
			"Failed to ping container with id `%d` and ip address `%s`: %v",
			container.GetId(), container.GetIpAddress(), err,
		)

		return
	}

	p.writeResult(container.GetId(), pinger.Statistics().AvgRtt.Milliseconds(), true)
}

func (p *GoPinger) writeResult(containerID int64, pingTime int64, wasSuccessful bool) {
	slot := p.tail.Add(1)
	p.slice[slot] = model.Ping{
		ContainerId:   containerID,
		PingTime:      pingTime,
		Date:          timestamppb.Now(),
		WasSuccessful: wasSuccessful,
	}
}
