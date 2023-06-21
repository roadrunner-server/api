package lock

import pq "github.com/roadrunner-server/api/v4/plugins/v2/priority_queue"

// Queue represents Lock plugin queue with it's elements types inside
type Queue interface {
	// Remove removes element with provided ID (if exists) and returns that elements
	Remove(id string) []pq.Item
	// Insert adds an item to the queue
	Insert(item pq.Item)
	// ExtractMin returns the item with the highest priority (less value is the highest priority)
	ExtractMin() pq.Item
	// Len returns the number of items in the queue
	Len() uint64
}
