package akka

type ActorRefScope interface {
	IsLocal() bool
}

type ActorRef interface {
	CanTell
	Path() ActorPath
	Equals(that interface{}) bool
}

type ActorRefWithCell interface {
	ActorRef

	Underlying() Cell
	Children() []ActorRef
	GetSingleChild(name string) InternalActorRef
}

type LocalActorRef interface {
	ActorRefWithCell
}

type NoSender struct {
}

func (NoSender) Path() (path ActorPath) {
	return
}
func (NoSender) Tell(message interface{}, sender ActorRef) {
	return
}
func (NoSender) Forward(message interface{}) {
	return
}
func (NoSender) Equals(that interface{}) bool {
	return false
}
func (NoSender) String() string {
	return ""
}

type LocalRef struct {
}

func (p *LocalRef) IsLocal() bool {
	return true
}
