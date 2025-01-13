module github.com/roadrunner-server/api/v4

go 1.23

toolchain go1.23.4

// Removed cmder API
retract v4.13.0

require (
	go.temporal.io/api v1.43.1
	go.uber.org/zap v1.27.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250106144421-5f5ef82da422
	google.golang.org/grpc v1.69.4
	google.golang.org/protobuf v1.36.2
)

require (
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250106144421-5f5ef82da422 // indirect
)
