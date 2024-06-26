package jobs

import (
	"context"
)

// Driver represents the interface for a single jobs driver
type Driver interface {
	// Push pushes the job to the underlying driver
	Push(ctx context.Context, msg Message) error
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

// Constructor constructs Consumer interface. Endure abstraction.
type Constructor interface {
	// Name returns the name of the driver
	Name() string
	// DriverFromConfig constructs a driver (e.g. kafka, amqp) from the configuration using the provided configKey
	DriverFromConfig(configKey string, queue Queue, pipeline Pipeline) (Driver, error)
	// DriverFromPipeline constructs a driver (e.g. kafka, amqp) from the pipeline. All configuration is provided by the pipeline
	DriverFromPipeline(pipe Pipeline, queue Queue) (Driver, error)
}
