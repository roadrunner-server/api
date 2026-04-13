module github.com/roadrunner-server/api/v4

go 1.25.0

// Removed cmder API
retract v4.13.0

require (
	go.temporal.io/api v1.62.8
	go.uber.org/zap v1.27.1
	google.golang.org/genproto/googleapis/api v0.0.0-20260406210006-6f92a3bedf2d
	google.golang.org/grpc v1.80.0
	google.golang.org/protobuf v1.36.11
)

require (
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.53.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260406210006-6f92a3bedf2d // indirect
)
