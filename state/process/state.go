package process

import (
	"fmt"
)

// Stater represents a worker state [BETA, internal]
type Stater interface {
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

// State provides information about specific worker.
type State struct {
	// Pid contains process id.
	Pid int64 `json:"pid"`

	// Status of the worker.
	Status int64 `json:"status"`

	// Number of worker executions.
	NumExecs uint64 `json:"numExecs"`

	// Created is unix nano timestamp of worker creation time.
	Created int64 `json:"created"`

	// MemoryUsage holds the information about worker memory usage in bytes.
	// Values might vary for different operating systems and based on RSS.
	MemoryUsage uint64 `json:"memoryUsage"`

	// CPU_Percent returns how many percent of the CPU time this process uses
	CPUPercent float64 `json:"CPUPercent"`

	// Command used in the service plugin and shows a command for the particular service
	Command string `json:"command"`

	// Status
	StatusStr string `json:"statusStr"`
}
