package go_queue

import (
	"github.com/owles/go-queue/contract"
	"github.com/owles/go-queue/driver/synced"
	"testing"
	"time"
)

// TestTask - Payload for test cases
type TestTask struct{}

func (receiver *TestTask) Signature() string {
	return "test_task"
}

func (receiver *TestTask) Frequency() *time.Duration {
	return nil
}

func (receiver *TestTask) Handle(args []contract.Arg) error {
	return nil
}

func TestQueue_Register(t *testing.T) {
	q := NewQueue(synced.NewDriver())
	q.Register(&TestTask{})

	if !q.Exist("test_task") {
		t.Fail()
	}

	err := q.Job(&TestTask{}, nil).Dispatch()
	if err != nil {
		t.Fail()
	}
}
