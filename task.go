package queue

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/owles/go-weby/contracts/queue"

	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/tasks"
)

type Task struct {
	connections *Connections
	connection  string
	chain       bool
	delay       *time.Time
	machinery   *Machinery
	jobs        []queue.Jobs
	queue       string
	server      *machinery.Server
}

func NewTask(connections *Connections, log *slog.Logger, job queue.Job, args []queue.Arg) *Task {
	return &Task{
		connections: connections,
		connection:  connections.GetDefault(),
		machinery:   NewMachinery(connections, log),
		jobs: []queue.Jobs{
			{
				Job:  job,
				Args: args,
			},
		},
	}
}

func NewChainTask(connections *Connections, log *slog.Logger, jobs []queue.Jobs) *Task {
	return &Task{
		connections: connections,
		connection:  connections.GetDefault(),
		chain:       true,
		machinery:   NewMachinery(connections, log),
		jobs:        jobs,
	}
}

func (receiver *Task) Delay(delay time.Time) queue.Task {
	receiver.delay = &delay

	return receiver
}

func (receiver *Task) Dispatch() error {
	conn := receiver.connections.Get(receiver.connection)
	if conn == nil {
		return fmt.Errorf("cannot find default connection")
	}

	if conn.Driver == DriverSync {
		return receiver.DispatchSync()
	}

	server, err := receiver.machinery.Server(receiver.connection, receiver.queue)
	if err != nil {
		return err
	}

	receiver.server = server

	if receiver.chain {
		return receiver.handleChain(receiver.jobs)
	} else {
		job := receiver.jobs[0]

		return receiver.handleAsync(job.Job, job.Args)
	}
}

func (receiver *Task) DispatchSync() error {
	if receiver.chain {
		for _, job := range receiver.jobs {
			if err := receiver.handleSync(job.Job, job.Args); err != nil {
				return err
			}
		}

		return nil
	} else {
		job := receiver.jobs[0]

		return receiver.handleSync(job.Job, job.Args)
	}
}

func (receiver *Task) OnConnection(connection string) queue.Task {
	receiver.connection = connection

	return receiver
}

func (receiver *Task) OnQueue(queue string) queue.Task {
	receiver.queue = queue

	return receiver
}

func (receiver *Task) handleChain(jobs []queue.Jobs) error {
	var signatures []*tasks.Signature
	for _, job := range jobs {
		var realArgs []tasks.Arg
		for _, arg := range job.Args {
			realArgs = append(realArgs, tasks.Arg{
				Type:  arg.Type,
				Value: arg.Value,
			})
		}

		signatures = append(signatures, &tasks.Signature{
			Name: job.Job.Signature(),
			Args: realArgs,
			ETA:  receiver.delay,
		})
	}

	chain, err := tasks.NewChain(signatures...)
	if err != nil {
		return err
	}

	_, err = receiver.server.SendChain(chain)

	return err
}

func (receiver *Task) handleAsync(job queue.Job, args []queue.Arg) error {
	var realArgs []tasks.Arg
	for _, arg := range args {
		realArgs = append(realArgs, tasks.Arg{
			Type:  arg.Type,
			Value: arg.Value,
		})
	}

	_, err := receiver.server.SendTask(&tasks.Signature{
		Name: job.Signature(),
		Args: realArgs,
		ETA:  receiver.delay,
	})
	if err != nil {
		return err
	}

	return nil
}

func (receiver *Task) handleSync(job queue.Job, args []queue.Arg) error {
	var realArgs []any
	for _, arg := range args {
		realArgs = append(realArgs, arg.Value)
	}

	return job.Handle(realArgs...)
}
