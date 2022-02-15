package worker

import (
	"context"
)

// Watcher is an interface for the Sync workers lifecycle
type Watcher interface {
	// Watch used to add workers to the container
	Watch(workers []BaseProcess) error

	// Take takes the first free worker
	Take(ctx context.Context) (BaseProcess, error)

	// Release releases the worker putting it back to the queue
	Release(w BaseProcess)

	// Allocate - allocates new worker and put it into the WorkerWatcher
	Allocate() error

	// Destroy destroys the underlying container
	Destroy(ctx context.Context)

	// Reset will replace container and workers array, kill all workers
	Reset(ctx context.Context)

	// List return all container w/o removing it from internal storage
	List() []BaseProcess

	// Remove will remove worker from the container
	Remove(wb BaseProcess)
}
