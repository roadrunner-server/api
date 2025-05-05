module github.com/roadrunner-server/api/v4

go 1.24

toolchain go1.24.1

// Removed cmder API
retract v4.13.0

require (
	go.temporal.io/api v1.49.0
	go.uber.org/zap v1.27.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250428153025-10db94c68c34
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.6
)

require (
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250428153025-10db94c68c34 // indirect
)
