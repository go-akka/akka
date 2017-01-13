package akka

type InternalActorRef interface {
	ActorRefScope
	ActorRef

	Provider() ActorRefProvider

	String() string
	Parent() InternalActorRef
	Child(names ...string) InternalActorRef

	Start()
	Resume(err error)
	Suspend()
	Restart(err error)
	Stop()
}
