package contract

import "time"

type Task interface {
	Dispatch() error
	Delay(time time.Duration) Task
	OnQueue(queue string) Task
}
