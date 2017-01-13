package akka

import (
	"reflect"
)

type MessageQueue interface {
	CleanUp(owner ActorRef, deadLetters MessageQueue) (err error)
	Dequeue() (envelope Envelope, err error)
	Enqueue(receiver ActorRef, handle Envelope) (err error)
	HasMessages() bool
	NumberOfMessages() int
}

type QueueBasedMessageQueue interface {
	MessageQueue
	Queue() []Envelope
}

type RequiresMessageQueue interface {
	RequiresMessageQueue() reflect.Type
}
