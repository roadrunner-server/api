package pool

type Queuer interface {
	// QueueSize can be implemented on the pool to provide the requests queue information
	QueueSize() uint64
}
