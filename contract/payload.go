package contract

import (
	"time"
)

type Payload interface {
	Uuid() string
	Fire()
	Release(delay *time.Duration)
	Delete()
	Attempts() int
	AvailableAt() time.Time
	GetSignature() string
	GetArgs() []Arg
}
