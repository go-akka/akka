package akka

import (
	"fmt"
)

var (
	_ InternalActorRef = (*MinimalActorRef)(nil)
)

var (
	NoBody *noBodyActorRef = newNoBodyActorRef()
)

type noBodyActorRef struct {
	*MinimalActorRef
	path ActorPath
}

func newNoBodyActorRef() *noBodyActorRef {
	path := NewRootActorPath(NewAddress("akka", "all-systems", "", 0), "/NoBody")

	return &noBodyActorRef{
		MinimalActorRef: NewMinimalActorRef(path, nil),
		path:            path,
	}
}

func (p *noBodyActorRef) Path() ActorPath {
	return p.path
}

func (p *noBodyActorRef) Provider() ActorRefProvider {
	panic("Nobody does not provide")
}

func (p *noBodyActorRef) CompareTo(other ActorRef) int {
	x := p.path.CompareTo(other.Path())
	if x == 0 {
		if p.path.Uid() < other.Path().Uid() {
			return -1
		} else if p.path.Uid() == other.Path().Uid() {
			return 0
		}
		return 1
	}
	return x
}

type MinimalActorRef struct {
	LocalRef

	path     ActorPath
	provider ActorRefProvider
}

func NewMinimalActorRef(path ActorPath, provider ActorRefProvider) *MinimalActorRef {
	return &MinimalActorRef{
		path:     path,
		provider: provider,
	}
}

func (p *MinimalActorRef) Tell(message interface{}, sender ...ActorRef) error {
	return nil
}

func (p *MinimalActorRef) Path() ActorPath {
	return p.path
}

func (p *MinimalActorRef) CompareTo(other ActorRef) int {
	x := p.path.CompareTo(other.Path())
	if x == 0 {
		if p.path.Uid() < other.Path().Uid() {
			return -1
		} else if p.path.Uid() == other.Path().Uid() {
			return 0
		}
		return 1
	}
	return x
}

func (p *MinimalActorRef) Provider() ActorRefProvider {
	return p.provider
}

func (p *MinimalActorRef) String() string {
	if p.path.Uid() == 0 {
		return fmt.Sprintf("Actor[%s]", p.path.String())
	}
	return fmt.Sprintf("Actor[%s]#[%d]", p.path.String(), p.path.Uid())
}

func (p *MinimalActorRef) Parent() InternalActorRef {
	return NoBody
}

func (p *MinimalActorRef) GetChild(names ...string) InternalActorRef {
	emptyCount := 0
	for i := 0; i < len(names); i++ {
		if names[i] == "" {
			emptyCount++
		}
	}

	if emptyCount == len(names) {
		return p
	}

	return NoBody
}

func (p *MinimalActorRef) Start() {
	return
}

func (p *MinimalActorRef) Resume(err error) {
	return
}

func (p *MinimalActorRef) Suspend() {
	return
}

func (p *MinimalActorRef) Restart(err error) {
	return
}

func (p *MinimalActorRef) Stop() {
	return
}

func (p *MinimalActorRef) SendSystemMessage(message SystemMessage) error {
	return nil
}
