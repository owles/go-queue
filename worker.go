package go_queue

import (
	"errors"
	"github.com/owles/go-queue/contract"
	"sync"
	"time"
)

type Worker struct {
	driver     contract.Driver
	queue      string
	concurrent int
}

func NewWorker(driver contract.Driver, queue string, concurrent int) *Worker {
	return &Worker{
		driver:     driver,
		queue:      queue,
		concurrent: concurrent,
	}
}

func (w *Worker) Run() error {
	if w.concurrent <= 0 {
		return errors.New("invalid number of concurrent workers")
	}

	var wg sync.WaitGroup
	errChan := make(chan error, w.concurrent)

	for i := 0; i < w.concurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				payload, err := w.driver.Pop(w.queue)
				if err != nil {
					errChan <- err
					return
				}

				if payload != nil {
					payload.Fire()
				}

				time.Sleep(time.Millisecond)
			}
		}()
	}

	wg.Wait()
	close(errChan)

	// Collect any errors that might have occurred during processing
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
