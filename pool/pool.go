package pool

import (
	"context"

	"github.com/roadrunner-server/api/v2/payload"
	"github.com/roadrunner-server/api/v2/worker"
)

// Pool managed set of inner worker processes.
type Pool interface {
	// GetConfig returns pool configuration.
	GetConfig() interface{}

	// Exec executes task with payload
	Exec(rqs *payload.Payload) (*payload.Payload, error)

	// ExecWithTTL executes task with context which is used with timeout
	ExecWithTTL(ctx context.Context, rqs *payload.Payload) (*payload.Payload, error)

	// Workers returns worker list associated with the pool.
	Workers() (workers []worker.BaseProcess)

	// RemoveWorker removes worker from the pool.
	RemoveWorker(worker worker.BaseProcess) error

	// Reset kill all workers inside the watcher and replaces with new
	Reset(ctx context.Context) error

	// Destroy all underlying stack (but let them to complete the task).
	Destroy(ctx context.Context)
}

// Streamer managed set of inner worker processes.
type Streamer interface {
	// ExecStream executes task with payload
	ExecStream(rqs *payload.Payload, resp chan *payload.Payload) error

	// ExecStreamWithTTL executes task with context which is used with timeout
	ExecStreamWithTTL(ctx context.Context, rqs *payload.Payload, resp chan *payload.Payload) error
}
