package priorityqueue

// Queue is a binary heap interface
type Queue interface {
	// Remove removes element with provided ID (if exists) and returns that elements
	Remove(id string) []Item
	// PeekPriority returns the highest priority
	PeekPriority() int64
	// Insert adds an item to the queue
	Insert(item Item)
	// ExtractMin returns the item with the highest priority (less value is the highest priority)
	ExtractMin() Item
	// Len returns the number of items in the queue
	Len() uint64
}

// Item represents a binary heap item
type Item interface {
	// ID returns a unique identifier for the item
	ID() string
	// Priority returns the priority level used to sort the item
	Priority() int64
	// Metadata returns the metadata for the item
	Metadata() map[string][]string
	// Body returns the payload associated with the item
	Body() []byte
	// Context returns any meta-information associated with the item
	Context() ([]byte, error)
}
