module github.com/roadrunner-server/api/v4

go 1.24

toolchain go1.24.2

// Removed cmder API
retract v4.13.0

require (
	go.temporal.io/api v1.52.0
	go.uber.org/zap v1.27.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250811230008-5f3141c8851a
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.7
)

require (
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250811230008-5f3141c8851a // indirect
)
