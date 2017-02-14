package akka

type ActorRefProvider interface {
	ActorOf(
		system ActorSystem,
		props Props,
		supervisor InternalActorRef,
		path ActorPath,
		systemService bool,
		deploy *Deploy,
		lookupDeploy bool,
		async bool) InternalActorRef

	DeadLetters() ActorRef
	Deployer() Deployer
	DefaultAddress() Address
	ExternalAddressFor(addr Address) Address
	Guardian() LocalActorRef
	Init(system ActorSystem) error

	RegisterTempActor(actorRef InternalActorRef, path ActorPath)
	ResolveActorRef(path ActorPath) ActorRef

	RootGuardian() InternalActorRef
	RootGuardianAt(address Address) ActorRef
	RootPath() ActorPath
	Settings() *Settings

	SystemGuardian() LocalActorRef
	TempContainer() InternalActorRef
	TempPath() ActorPath
	TerminationFuture()
	UnregisterTempActor(path ActorPath)
}

type LocalActorRefProvider interface {
	ActorRefProvider
	LocalActorRefProvider()
}

type ActorRefFactory interface {
	ActorOf(props Props, name string) (ref ActorRef, err error)
	ActorSelection(path ActorPath) (selection ActorSelection, err error)
}
