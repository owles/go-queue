package synced

import (
	"github.com/owles/go-queue/contract"
	"sync"
	"time"
)

type Driver struct {
	queue map[string][]contract.Payload
	mutex sync.Mutex
}

func NewDriver() *Driver {
	return &Driver{
		queue: make(map[string][]contract.Payload),
	}
}

func (receiver *Driver) Size(queue string) int {
	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()
	return len(receiver.queue[queue])
}

func (receiver *Driver) Push(payload contract.Payload, queue string) error {
	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()
	return nil
}

func (receiver *Driver) Pop(queue string) (contract.Payload, error) {
	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if len(receiver.queue[queue]) == 0 {
		return nil, nil
	}

	for i, payload := range receiver.queue[queue] {
		if time.Now().After(payload.AvailableAt()) {
			receiver.queue[queue] = append(receiver.queue[queue][:i], receiver.queue[queue][i+1:]...)
			return payload, nil
		}
	}

	return nil, nil
}
