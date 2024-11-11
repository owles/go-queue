package queue

import (
	"fmt"
	"github.com/owles/go-weby/contracts/queue"
)

func jobs2Tasks(jobs []queue.Job) (map[string]any, error) {
	tasks := make(map[string]any)

	for _, job := range jobs {
		if job.Signature() == "" {
			return nil, fmt.Errorf("empty Job signature")
		}

		if tasks[job.Signature()] != nil {
			return nil, fmt.Errorf("duplicate Job signature: %s", job.Signature())
		}

		tasks[job.Signature()] = job.Handle
	}

	return tasks, nil
}
