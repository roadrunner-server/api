module github.com/roadrunner-server/api/v4

go 1.24

toolchain go1.24.2

// Removed cmder API
retract v4.13.0

require (
	go.temporal.io/api v1.51.0
	go.uber.org/zap v1.27.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250721164621-a45f3dfb1074
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.6
)

require (
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250721164621-a45f3dfb1074 // indirect
)
