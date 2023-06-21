package jobs

import (
	pq "github.com/roadrunner-server/api/v4/plugins/v2/priority_queue"
)

// Queue represents JOBS plugin queue with it's elements types inside
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
	pq.Item
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
	// Headers returns the metadata for the item
	Headers() map[string][]string
}

// Message represents the protobuf message received from the RPC call
type Message interface {
	pq.Item
	KafkaOptions
	// Name returns the name of the Job
	Name() string
	// Payload returns the data associated with the job
	Payload() string
	// Delay returns the delay time for the Job (not supported by all drivers)
	Delay() int64
	// AutoAck returns the autocommit status for the Job
	AutoAck() bool
	// UpdatePriority sets the priority of the Job. Priority is optional but cannot be set to 0.
	// The default priority is 10
	UpdatePriority(int64)
	// Headers returns the metadata for the item
	Headers() map[string][]string
}

// KAFKA options (leave them empty for other drivers)
type KafkaOptions interface {
	// Offset returns the offset associated with the Job
	Offset() int64
	// Partition returns the partition associated with the Job
	Partition() int32
	// Topic returns the topic associated with the Job
	Topic() string
	// Metadata returns the metadata associated with the Job
	Metadata() string
}

type Pipeline interface {
	// With sets a pipeline value
	With(name string, value interface{})
	// Name returns the pipeline name.
	Name() string
	// Driver returns the driver associated with the pipeline.
	Driver() string
	// Has checks if a value is present in the pipeline.
	Has(name string) bool
	// String returns the value of an option as a string or the default value.
	String(name string, d string) string
	// Int returns the value of an option as an int or the default value.
	Int(name string, d int) int
	// Bool returns the value of an option as a bool or the default value.
	Bool(name string, d bool) bool
	// Map returns the nested map value or an empty config.
	// This might be used for SQS attributes or tags, for example
	Map(name string, out map[string]string) error
	// Priority returns the default pipeline priority
	Priority() int64
	// Get is used to retrieve the data associated with a key
	Get(key string) interface{}
}
