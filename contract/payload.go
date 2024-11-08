package contract

import (
	"time"
)

type Payload interface {
	Fire()
	Release(delay *time.Duration)
	Delete()
	Attempts() int
	AvailableAt() time.Time
	GetSignature() string
	GetArgs() []Arg
}
