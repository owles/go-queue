package contract

type Driver interface {
	Size(queue string) int
	Push(payload Payload, queue string) error
	Pop(queue string) (Payload, error)
}
