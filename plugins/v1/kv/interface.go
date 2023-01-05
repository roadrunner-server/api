package kv

import kvv1 "go.buf.build/protocolbuffers/go/roadrunner-server/api/kv/v1"

// Storage represents single abstract storage.
type Storage interface {
	// Has checks if value exists.
	Has(keys ...string) (map[string]bool, error)

	// Get loads value content into a byte slice.
	Get(key string) ([]byte, error)

	// MGet loads content of multiple values
	// Returns the map with existing keys and associated values
	MGet(keys ...string) (map[string][]byte, error)

	// Set used to upload item to KV with TTL
	// 0 value in TTL means no TTL
	Set(items ...*kvv1.Item) error

	// MExpire sets the TTL for multiply keys
	MExpire(items ...*kvv1.Item) error

	// TTL return the rest time to live for provided keys
	// Not supported for the memcached
	TTL(keys ...string) (map[string]string, error)

	// Clear clean the entire storage
	Clear() error

	// Delete one or multiple keys.
	Delete(keys ...string) error

	// Stop the storage driver
	Stop()
}

// Constructor provides storage based on the config
type Constructor interface {
	// KvFromConfig provides Storage based on the config key
	KvFromConfig(key string) (Storage, error)
	Name() string
}
