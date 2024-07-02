package jobs

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
	// Active jobs which are consumed by the driver but not handled by the PHP worker yet
	Active int64
	// Delayed jobs
	Delayed int64
	// Reserved jobs which are in the driver but not consumed yet
	Reserved int64
	// Status - 1 Ready, 0 - Paused
	Ready bool
	// New in 2.10.5, pipeline priority
	Priority uint64
	// ErrorMessage New in 2023.1
	ErrorMessage string
}
