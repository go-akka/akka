package akka

type ChildStats interface {
	ChildStats()
}

type ChildNameReserved interface {
	ChildStats

	ChildNameReserved()
}

type ChildRestartStats interface {
	ChildStats

	Child() InternalActorRef
	ChildRestartStats()
}

type ChildrenContainer interface {
	Add(name string, stats ChildRestartStats) ChildrenContainer
	Remove(child ActorRef) ChildrenContainer

	GetByName(name string) (stats ChildStats, exist bool)
	GetByRef(actor ActorRef) (stats ChildRestartStats, exist bool)

	Children() []ActorRef
	Stats() []ChildRestartStats

	ShallDie(actor ActorRef) ChildrenContainer
	Reserve(name string) ChildrenContainer
	Unreserve(name string) ChildrenContainer

	IsTerminating() bool
	IsNormal() bool
	Contains(actor ActorRef) bool
}
