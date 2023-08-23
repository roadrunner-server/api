package priority_queue

// Item interface represents the base meta-information which any priority queue message must have
type Item interface {
	// ID returns a unique identifier for the item
	ID() string
	// GroupID returns the group associated with the item, used to remove all items with the same groupID
	GroupID() string
	// Priority returns the priority level used to sort the item
	Priority() int64
}
