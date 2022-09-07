package process

// State represents a worker state [BETA, internal]
type State interface {
	Pid() int
	Status() string
	NumJobs() uint64
	Created() int64
	MemoryUsage() uint64
	CPUPercent() float64
	Command() string
}
