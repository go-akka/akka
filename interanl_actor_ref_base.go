package akka

type InternalActorRefBase struct {
	parent *InternalActorRefBase
	child  *InternalActorRefBase
}

func (p *InternalActorRefBase) Path() (path ActorPath) {
	return
}

func (p *InternalActorRefBase) Tell(message interface{}, sender ...ActorRef) error {
	return nil
}

func (p *InternalActorRefBase) Forward(message interface{}) {
	return
}

func (p *InternalActorRefBase) CompareTo(other ActorRef) int {
	return 0
}

func (p *InternalActorRefBase) String() string {
	return ""
}

func (p *InternalActorRefBase) Parent() ActorRef {
	return p.parent
}

func (p *InternalActorRefBase) Child(names ...string) *InternalActorRefBase {
	return p.child
}

func (p *InternalActorRefBase) IsLocal() bool {
	return true
}

func (p *InternalActorRefBase) isTerminated() bool {
	return false
}
