module github.com/roadrunner-server/api/v4

go 1.22.4

// Removed cmder API
retract v4.13.0

require (
	go.temporal.io/api v1.34.0
	go.uber.org/zap v1.27.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240624140628-dc46fd24d27d
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
)

require (
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240624140628-dc46fd24d27d // indirect
)
