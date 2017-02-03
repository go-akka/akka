package akka

import (
	"time"
)

type ReceiveFunc func(message interface{}) (handled bool, err error)

type CanWatch interface {
	Watch(subject ActorRef) (err error)
	Unwatch(subject ActorRef)
}

type ActorContext interface {
	ActorRefFactory
	CanWatch

	Become(receive ReceiveFunc, discardOld bool) (err error)
	Unbecome()

	Child(name string) (ref ActorRef, exist bool)
	Children() (refs []ActorRef)
	Parent() (ref ActorRef)

	Props() (props Props)

	ReceiveTimeout() (timeout time.Duration)
	SetReceiveTimeout(timeout time.Duration)

	Self() ActorRef
	Sender() ActorRef

	System() ActorSystem

	StopChild(actor ActorRef)
}
