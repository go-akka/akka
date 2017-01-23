package akka

type ActorRefScope interface {
	IsLocal() bool
}

type ActorRef interface {
	CanTell
	Path() ActorPath
	CompareTo(other ActorRef) int
}

type ActorRefWithCell interface {
	InternalActorRef

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
func (NoSender) Tell(message interface{}, sender ...ActorRef) error {
	return nil
}
func (NoSender) Forward(message interface{}) {
	return
}
func (NoSender) CompareTo(other ActorRef) int {
	return 0
}
func (NoSender) String() string {
	return ""
}

type LocalRef struct {
}

func (p *LocalRef) IsLocal() bool {
	return true
}
