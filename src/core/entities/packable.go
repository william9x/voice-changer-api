package entities

type Task interface {
	Packable
	QueueOpt
}

type Packable interface {
	Pack() ([]byte, error)
}

type QueueOpt interface {
	Queue() string
}
