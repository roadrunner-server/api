package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/roadrunner-server/api/v2/payload"
	"github.com/roadrunner-server/goridge/v3/pkg/relay"
)

// Allocator is responsible for worker allocation in the pool
type Allocator func() (BaseProcess, error)

// Streamer interface adds streams capabilities to the pool
// [BETA] interface, might be changed later
type Streamer interface {
	// BaseProcess provides basic functionality for the SyncWorker
	BaseProcess
	// ExecStream used to execute payload on the SyncWorker, there is no TIMEOUTS
	ExecStream(rqs *payload.Payload, resp chan *payload.Payload, stopCh chan struct{}) error
	// ExecStreamWithTTL used to handle Exec with TTL
	ExecStreamWithTTL(ctx context.Context, p *payload.Payload, resp chan *payload.Payload, stopCh chan struct{}) error
}

// Worker is a non-bc replacement for the SDK
type Worker = SyncWorker

// SyncWorker is not a good name, since it's just a sync executor, but not all worker is sync
type SyncWorker interface {
	// BaseProcess provides basic functionality for the SyncWorker
	BaseProcess
	// Exec used to execute payload on the SyncWorker, there is no TIMEOUTS
	Exec(rqs *payload.Payload) (*payload.Payload, error)
	// ExecWithTTL used to handle Exec with TTL
	ExecWithTTL(ctx context.Context, p *payload.Payload) (*payload.Payload, error)
}

// State represents WorkerProcess status and updated time.
type State interface {
	fmt.Stringer
	// Value returns StateImpl value
	Value() int64
	// Set sets the StateImpl
	Set(value int64)
	// NumExecs shows how many times WorkerProcess was invoked
	NumExecs() uint64
	// IsActive returns true if WorkerProcess not Inactive or Stopped
	IsActive() bool
	// RegisterExec using to registering php executions
	RegisterExec()
	// SetLastUsed sets worker last used time
	SetLastUsed(lu uint64)
	// LastUsed return worker last used time
	LastUsed() uint64
}

type BaseProcess interface {
	fmt.Stringer

	// Pid returns worker pid.
	Pid() int64

	// Created returns time, worker was created at.
	Created() time.Time

	// State return receive-only WorkerProcess state object, state can be used to safely access
	// WorkerProcess status, time when status changed and number of WorkerProcess executions.
	State() FSM

	// Start used to run Cmd and immediately return
	Start() error

	// Wait must be called once for each WorkerProcess, call will be released once WorkerProcess is
	// complete and will return process error (if any), if stderr is presented it is value
	// will be wrapped as WorkerError. Method will return error code if php process fails
	// to find or Start the script.
	Wait() error

	// Stop sends soft termination command to the WorkerProcess and waits for process completion.
	Stop() error

	// Kill kills underlying process, make sure to call Wait() func to gather
	// error log from the stderr. Does not wait for process completion!
	Kill() error

	// Relay returns attached to worker goridge relay
	Relay() relay.Relay

	// AttachRelay used to attach goridge relay to the worker process
	AttachRelay(rl relay.Relay)
}
