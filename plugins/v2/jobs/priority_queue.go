package jobs

// Queue is a binary heap interface
type Queue interface {
	// Remove removes element with provided ID (if exists) and returns that elements
	Remove(id string) []Job
	// Insert adds an item to the queue
	Insert(item Job)
	// ExtractMin returns the item with the highest priority (less value is the highest priority)
	ExtractMin() Job
	// Len returns the number of items in the queue
	Len() uint64
}

// Job represents a binary heap item
type Job interface {
	Base
	// Ack acknowledges the item after processing
	Ack() error
	// Nack discards the item
	Nack() error
	// Requeue puts the message back to the queue with an optional delay
	Requeue(headers map[string][]string, delay int64) error
	// Body returns the payload associated with the item
	Body() []byte
	// Context returns any meta-information associated with the item
	Context() ([]byte, error)
}

// Base interface represents the base meta-information which any message must have
type Base interface {
	// ID returns a unique identifier for the item
	ID() string
	// PipelineID returns the pipeline ID associated with the item
	PipelineID() string
	// Priority returns the priority level used to sort the item
	Priority() int64
	// Headers returns the metadata for the item
	Headers() map[string][]string
}
