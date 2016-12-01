package actor

var (
	_ ActorRef = (*ActorRefBase)(nil)
)

type ActorRefScope interface {
	IsLocal() bool
}

type ActorRef interface {
	Path() ActorPath
	Tell(message interface{}, sender ActorRef)
	Forward(message interface{})
	Equals(that interface{}) bool
	String() string
}

type ActorRefBase struct {
	path    ActorPath
	context ActorContext
}

func (p *ActorRefBase) Path() ActorPath {
	return p.path
}

func (p *ActorRefBase) Tell(message interface{}, sender ActorRef) {
	return
}

func (p *ActorRefBase) Forward(message interface{}) {
	return
}

func (p *ActorRefBase) Equals(that interface{}) bool {
	switch other := that.(type) {
	case ActorPath:
		{
			return p.path.Equals(other)
		}
	}
	return false
}

func (p *ActorRefBase) String() string {
	return ""
}
