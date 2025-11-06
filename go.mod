module github.com/roadrunner-server/api/v4

go 1.24.0

toolchain go1.24.2

// Removed cmder API
retract v4.13.0

require (
	go.temporal.io/api v1.57.0
	go.uber.org/zap v1.27.0
	google.golang.org/genproto/googleapis/api v0.0.0-20251103181224-f26f9409b101
	google.golang.org/grpc v1.76.0
	google.golang.org/protobuf v1.36.10
)

require (
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251103181224-f26f9409b101 // indirect
)
