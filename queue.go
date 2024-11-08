package go_queue

import (
	"github.com/owles/go-queue/contract"
	"sync"
)

type Queue struct {
	driver contract.Driver
	jobs   sync.Map
}

func NewQueue(driver contract.Driver) *Queue {
	return &Queue{
		driver: driver,
	}
}

// Exist - check Payload Signature in the Queue Jobs
func (receiver *Queue) Exist(signature string) bool {
	_, ok := receiver.jobs.Load(signature)
	return ok
}

// Register - Add Payload Struct in the Queue
func (receiver *Queue) Register(task contract.Job) {
	receiver.jobs.LoadOrStore(task.Signature(), task)
}

func (receiver *Queue) Job(job contract.Job, args []contract.Arg) contract.Task {
	return NewTask(
		receiver.driver,
		job,
		args,
	)
}
