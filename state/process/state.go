package process

// State provides information about specific worker.
type State struct {
	// Pid contains process id.
	Pid int `json:"pid"`

	// Status of the worker.
	Status string `json:"status"`

	// Number of worker executions.
	NumJobs uint64 `json:"numExecs"`

	// Created is unix nano timestamp of worker creation time.
	Created int64 `json:"created"`

	// MemoryUsage holds the information about worker memory usage in bytes.
	// Values might vary for different operating systems and based on RSS.
	MemoryUsage uint64 `json:"memoryUsage"`

	// CPU_Percent returns how many percent of the CPU time this process uses
	CPUPercent float64

	// Command used in the service plugin and shows a command for the particular service
	Command string
}
