package config

import (
	"time"
)

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshals it into a Struct.
	//
	// func (h *HttpService) Init(cp config.Configurer) error {
	//     h.config := &HttpConfig{}
	//     if err := configProvider.UnmarshalKey("http", h.config); err != nil {
	//         return err
	//     }
	// }
	UnmarshalKey(name string, out any) error

	// Unmarshal unmarshal the config into a Struct. Make sure that the tags
	// on the fields of the structure are properly set.
	Unmarshal(out any) error

	// Get used to get config section
	Get(name string) any

	// Overwrite used to overwrite particular values in the unmarshalled config
	Overwrite(values map[string]any) error

	// Has checks if config section exists.
	Has(name string) bool

	// GracefulTimeout represents timeout for all servers registered in the endure
	GracefulTimeout() time.Duration

	// RRVersion returns running RR version
	RRVersion() string
}
