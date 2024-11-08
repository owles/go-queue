package go_queue

import (
	"github.com/owles/go-queue/contract"
	"time"
)

type Payload struct {
	driver contract.Driver

	signature string
	args      []contract.Arg

	attempts    int
	availableAt time.Time
}

func NewPayload(driver contract.Driver, signature string, availableAt time.Time, args []contract.Arg) *Payload {
	return &Payload{
		driver: driver,

		signature: signature,
		args:      args,

		attempts:    0,
		availableAt: availableAt,
	}
}

func (receiver *Payload) Fire() {
}

func (receiver *Payload) Release(delay *time.Duration) {
}

func (receiver *Payload) Delete() {
}

func (receiver *Payload) Attempts() int {
	return receiver.attempts
}

func (receiver *Payload) AvailableAt() time.Time {
	return receiver.availableAt
}

func (receiver *Payload) GetSignature() string {
	return receiver.signature
}

func (receiver *Payload) GetArgs() []contract.Arg {
	return receiver.args
}
