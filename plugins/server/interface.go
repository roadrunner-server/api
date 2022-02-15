package server

import (
	"context"
	"os/exec"

	"github.com/roadrunner-server/api/v2/pool"
	"github.com/roadrunner-server/api/v2/worker"
	"go.uber.org/zap"
)

// Server creates workers for the application.
type Server interface {
	// CmdFactory return a new command based on the .rr.yaml server.command section
	CmdFactory(env map[string]string) func() *exec.Cmd
	// NewWorker return a new worker with provided and attached by the user listeners and environment variables
	NewWorker(ctx context.Context, env map[string]string) (worker.BaseProcess, error)
	// NewWorkerPool return new pool of workers (PHP) with attached events listeners, env variables and based on the provided configuration
	NewWorkerPool(ctx context.Context, config interface{}, env map[string]string, log *zap.Logger) (pool.Pool, error)
}
