package event_bus //nolint:stylecheck

import (
	"fmt"
)

type EventBus interface {
	SubscribeAll(subID string, ch chan<- Event) error
	SubscribeP(subID string, pattern string, ch chan<- Event) error
	Unsubscribe(subID string)
	UnsubscribeP(subID, pattern string)
	Len() uint
	Send(ev Event)
}

type Event interface {
	Type() fmt.Stringer
	Plugin() string
	Message() string
}
