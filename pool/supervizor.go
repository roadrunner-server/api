package pool

type Supervised interface {
	Pool
	// Start used to start watching process for all pool workers
	Start()
}
