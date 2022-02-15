package worker

// SYNC WITH worker_watcher.GET
const (
	// StateInactive - no associated process
	StateInactive int64 = iota

	// StateReady - ready for job.
	StateReady

	// StateWorking - working on given payload.
	StateWorking

	// StateInvalid - indicates that WorkerProcess is being disabled and will be removed.
	StateInvalid

	// StateStopping - process is being softly stopped.
	StateStopping

	// StateKilling - process is being forcibly stopped
	StateKilling

	// StateDestroyed State of worker, when no need to allocate new one
	StateDestroyed

	// StateMaxJobsReached State of worker, when it reached executions limit
	StateMaxJobsReached

	// StateStopped - process has been terminated.
	StateStopped

	// StateErrored - error StateImpl (can't be used).
	StateErrored
)
