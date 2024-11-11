package contract

type Worker interface {
	Run() error
}

type Args struct {
	Driver     Driver
	Queue      string
	Concurrent int
}
