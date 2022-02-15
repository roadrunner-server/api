package payload

import (
	"github.com/roadrunner-server/api/v2/internal"
)

// Payload carries binary header and body to stack and
// back to the server.
type Payload struct {
	// Context represent payload context, might be omitted.
	Context []byte

	// body contains binary payload to be processed by WorkerProcess.
	Body []byte

	// Type of codec used to decode/encode payload
	Codec byte
}

// String returns payload body as string
func (p *Payload) String() string {
	return internal.AsString(p.Body)
}
