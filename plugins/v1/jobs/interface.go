package jobs

import (
	"context"

	"github.com/roadrunner-server/api/v3/plugins/v1/priority_queue"
	"go.uber.org/zap"
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

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshals it into a Struct.
	UnmarshalKey(name string, out any) error

	// Has checks if config section exists.
	Has(name string) bool
}

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

type Job interface {
	Name() string
	ID() string
	Payload() string
	Headers() map[string][]string

	Pipeline() string
	Priority() int64
	Delay() int64
	AutoAck() bool

	SetPriority(int64)
}

type Pipeline interface {
	// With pipeline value
	With(name string, value any)
	// Name returns pipeline name.
	Name() string
	// Driver associated with the pipeline.
	Driver() string
	// Has checks if value presented in pipeline.
	Has(name string) bool
	// String must return option value as string or return default value.
	String(name string, d string) string
	// Int must return option value as string or return default value.
	Int(name string, d int) int
	// Bool must return option value as bool or return default value.
	Bool(name string, d bool) bool
	// Map must return nested map value or empty config.
	// Here might be sqs attributes or tags for example
	Map(name string, out map[string]string) error
	// Priority returns default pipeline priority
	Priority() int64
	// Get used to get the data associated with the key
	Get(key string) any
}

// Consumer represents a single jobs driver interface
type Consumer interface {
	Push(ctx context.Context, job Job) error
	Register(ctx context.Context, pipeline Pipeline) error
	Run(ctx context.Context, pipeline Pipeline) error
	Stop(ctx context.Context) error

	Pause(ctx context.Context, pipeline string)
	Resume(ctx context.Context, pipeline string)

	// State provide information about driver state
	State(ctx context.Context) (*State, error)
}

// Acknowledger provides queue specific item management
type Acknowledger interface {
	// Ack - acknowledge the Item after processing
	Ack() error

	// Nack - discard the Item
	Nack() error

	// Requeue - put the message back to the queue with the optional delay
	Requeue(headers map[string][]string, delay int64) error

	// Respond to the queue
	Respond(payload []byte, queue string) error
}

// Constructor constructs Consumer interface. Endure abstraction.
type Constructor interface {
	Name() string
	ConsumerFromConfig(configKey string, queue priorityqueue.Queue) (Consumer, error)
	ConsumerFromPipeline(pipe Pipeline, queue priorityqueue.Queue) (Consumer, error)
}
