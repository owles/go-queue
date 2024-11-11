package go_queue

import (
	"github.com/owles/go-queue/driver/synced"
	"testing"
	"time"
)

func TestQueue_Worker(t *testing.T) {
	q := NewQueue(synced.NewDriver())
	q.Register(&TestTask{})

	if !q.Exist("test_task") {
		t.Fail()
	}

	err := q.Job(&TestTask{}, nil).Dispatch()
	if err != nil {
		t.Fail()
	}

	go q.Worker().Run()

	time.Sleep(time.Second * 5)
}
