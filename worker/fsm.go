package worker

import (
	"fmt"
)

// FSM represents endure finite state machine
type FSM interface {
	fmt.Stringer
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
	// CurrentState returns the current state of the FSM
	CurrentState() int64
	// Transition used to move from one state to another
	Transition(from int64)
	// Compare compares state to the actual state and return true if they equal
	// false otherwise
	Compare(state int64) bool
}
