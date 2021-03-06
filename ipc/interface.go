package ipc

import (
	"context"
	"os/exec"

	"github.com/roadrunner-server/api/v2/worker"
)

// Factory is responsible for wrapping given command into tasks WorkerProcess.
type Factory interface {
	// SpawnWorkerWithTimeout creates new WorkerProcess process based on given command with context.
	// Process must not be started.
	SpawnWorkerWithTimeout(context.Context, *exec.Cmd) (worker.BaseProcess, error)
	// SpawnWorker creates new WorkerProcess process based on given command.
	// Process must not be started.
	SpawnWorker(*exec.Cmd) (worker.BaseProcess, error)
	// Close the factory and underlying connections.
	Close() error
}
