package process

import (
	"fmt"
)

// State represents a worker state [BETA, internal]
type State interface {
	fmt.Stringer
	// Pid contains process id
	Pid() int64
	// Status of the worker
	Status() int64
	// NumExecs is number of worker executions.
	NumExecs() uint64
	// Created is unix nano timestamp of worker creation time.
	Created() int64
	// MemoryUsage holds the information about worker memory usage in bytes.
	// Values might vary for different operating systems and based on RSS.
	MemoryUsage() uint64
	// CPUPercent returns how many percent of the CPU time this process uses
	CPUPercent() float64
	// Command used in the service plugin and shows a command for the particular service
	Command() string
}
