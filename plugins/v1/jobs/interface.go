package jobs

import (
	"context"

	pq "github.com/roadrunner-server/api/v3/plugins/v1/priority_queue"
)

// constant keys to pack/unpack messages from different drivers
const (
	RRID       string = "rr_id"
	RRJob      string = "rr_job"
	RRHeaders  string = "rr_headers"
	RRPipeline string = "rr_pipeline"
	RRDelay    string = "rr_delay"
	RRPriority string = "rr_priority"
	RRAutoAck  string = "rr_auto_ack"
)

type Command string

const (
	Stop Command = "stop"
)

// State represents job's state
type State struct {
	// Pipeline name
	Pipeline string
	// Driver name
	Driver string
	// Queue name (tube for the beanstalk)
	Queue string
	// Active jobs which are consumed from the driver but not handled by the PHP worker yet
	Active int64
	// Delayed jobs
	Delayed int64
	// Reserved jobs which are in the driver but not consumed yet
	Reserved int64
	// Status - 1 Ready, 0 - Paused
	Ready bool
	// New in 2.10.5, pipeline priority
	Priority uint64
}

// Job represents the unit of work that a user may push to the Jobs plugin
type Job interface {
	// Name returns the name of the Job
	Name() string
	// ID returns a unique identifier for the Job (e.g. UUID)
	ID() string
	// Payload returns the data associated with the job
	Payload() string
	// Headers returns the headers in the format of map[string][]string
	Headers() map[string][]string

	// Pipeline returns the pipeline associated with the job
	Pipeline() string
	// Priority returns the priority level of the Job
	Priority() int64
	// Delay returns the delay time for the Job (not supported by all drivers)
	Delay() int64
	// AutoAck returns the autocommit status for the Job
	AutoAck() bool

	// --KAFKA options (leave them empty for other drivers)

	// Offset returns the offset associated with the Job
	Offset() int64
	// Partition returns the partition associated with the Job
	Partition() int32
	// Topic returns the topic associated with the Job
	Topic() string
	// Metadata returns the metadata associated with the Job
	Metadata() string

	// UpdatePriority sets the priority of the Job. Priority is optional but cannot be set to 0.
	// The default priority is 10
	UpdatePriority(int64)
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

// Driver represents the interface for a single jobs driver
type Driver interface {
	// Push pushes the job to the underlying driver
	Push(ctx context.Context, job Job) error
	// Run starts consuming the pipeline
	Run(ctx context.Context, pipeline Pipeline) error
	// Stop stops the consumer and closes the underlying connection
	Stop(ctx context.Context) error
	// Pause pauses the jobs consuming (while still allowing job pushing)
	Pause(ctx context.Context, pipeline string) error
	// Resume resumes the consumer
	Resume(ctx context.Context, pipeline string) error
	// State returns information about the driver state
	State(ctx context.Context) (*State, error)
}

// Acknowledger provides queue-specific item management
type Acknowledger interface {
	// Ack acknowledges the item after processing
	Ack() error
	// Nack discards the item
	Nack() error
	// Requeue puts the message back to the queue with an optional delay
	Requeue(headers map[string][]string, delay int64) error
	// Respond sends a response to the queue
	Respond(payload []byte, queue string) error
}

// Commanders provides the ability to send a command to the Jobs plugin
type Commanders interface {
	// Command returns the command name
	Command() Command
	// Pipeline returns the associated command pipeline
	Pipeline() string
}

// Constructor constructs Consumer interface. Endure abstraction.
type Constructor interface {
	// Name returns the name of the driver
	Name() string
	// DriverFromConfig constructs a driver (e.g. kafka, amqp) from the configuration using the provided configKey
	DriverFromConfig(configKey string, queue pq.Queue, pipeline Pipeline, cmder chan<- Commanders) (Driver, error)
	// DriverFromPipeline constructs a driver (e.g. kafka, amqp) from the pipeline. All configuration is provided by the pipeline
	DriverFromPipeline(pipe Pipeline, queue pq.Queue, cmder chan<- Commanders) (Driver, error)
}
