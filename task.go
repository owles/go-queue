package go_queue

import (
	"github.com/owles/go-queue/contract"
	"time"
)

type Task struct {
	driver contract.Driver

	delay *time.Duration
	queue string

	job  contract.Job
	args []contract.Arg
}

func NewTask(driver contract.Driver, job contract.Job, args []contract.Arg) *Task {
	return &Task{
		driver: driver,

		queue: "",
		delay: nil,

		job:  job,
		args: args,
	}
}

func (receiver *Task) Delay(delay time.Duration) contract.Task {
	receiver.delay = &delay
	return receiver
}

func (receiver *Task) OnQueue(queue string) contract.Task {
	receiver.queue = queue
	return receiver
}

func (receiver *Task) Dispatch() error {
	availableAt := time.Now().UTC()
	if receiver.delay != nil {
		availableAt.Add(*receiver.delay)
	}

	return receiver.driver.Push(
		NewPayload(
			receiver.driver,
			receiver.job.Signature(),
			availableAt,
			receiver.args,
		),
		receiver.queue,
	)
}
