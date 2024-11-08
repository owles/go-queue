package contract

import "time"

type Job interface {
	Handle(args []Arg) error

	Signature() string
	Frequency() *time.Duration
}

type Arg struct {
	Type  string `bson:"type"`
	Value string `bson:"value"`
}
