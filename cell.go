package akka

type Cell interface {
	Self() ActorRef
	System() ActorSystem
	Dispatcher() MessageDispatcher
	AttachChild(props Props, name string, systemService bool) (ref ActorRef, err error)

	Start()
	Suspend()
	Resume(err error)
	Restart(err error)
	Stop() (err error)

	Parent() ActorRef
	IsLocal() bool
	Props() Props

	HasMessages() bool
	NumberOfMessages() int
	SendMessage(msg Envelope) (err error)

	IsTerminated() bool

	ChildrenRefs() ChildrenContainer
	GetSingleChild(name string) ActorRef
	GetChildByName(name string) (stats ChildStats, exist bool)
}
