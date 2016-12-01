package actor

type Receive func(message interface{}) bool

type ActorContext interface {
	Self() ActorRef
	Sender() ActorRef
	Dispatcher()
	System() ActorSystem
	Parent() ActorRef
	GetChildren() []ActorRef
	// Props()

	ActorOf(props Props, name string) (ref ActorRef, err error)

	Become(receive Receive) (err error)
	BecomeStacked(receive Receive) (err error)
	UnbecomeStacked() (err error)

	Watch(subject ActorRef)
	Unwatch(subject ActorRef)

	Stop(child ActorRef)
}
