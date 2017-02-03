package internal

import (
	"github.com/go-akka/akka"
)

var (
	EmptyChildrenContainerInstance = &EmptyChildrenContainer{emptyStats: NewImmutableMap()}
)

type EmptyChildrenContainer struct {
	emptyStats *ImmutableMap
}

func (p *EmptyChildrenContainer) Add(name string, stats akka.ChildRestartStats) akka.ChildrenContainer {
	return createNormalChildContainer(p.emptyStats.Add(name, stats))
}

func (p *EmptyChildrenContainer) Remove(child akka.ActorRef) akka.ChildrenContainer {
	return p
}

func (p *EmptyChildrenContainer) GetByName(name string) (stats akka.ChildStats, exist bool) {
	return nil, false
}

func (p *EmptyChildrenContainer) GetByRef(actor akka.ActorRef) (stats akka.ChildRestartStats, exist bool) {
	return nil, false
}

func (p *EmptyChildrenContainer) Children() []akka.ActorRef {
	return nil
}

func (p *EmptyChildrenContainer) Stats() []akka.ChildRestartStats {
	return nil
}

func (p *EmptyChildrenContainer) ShallDie(actor akka.ActorRef) akka.ChildrenContainer {
	return p
}

func (p *EmptyChildrenContainer) Reserve(name string) akka.ChildrenContainer {
	return createNormalChildContainer(p.emptyStats.Add(name, _childNameReservedInstance))
}

func (p *EmptyChildrenContainer) Unreserve(name string) akka.ChildrenContainer {
	return p
}

func (p *EmptyChildrenContainer) IsTerminating() bool {
	return false
}

func (p *EmptyChildrenContainer) IsNormal() bool {
	return true
}

func (p *EmptyChildrenContainer) Contains(actor akka.ActorRef) bool {
	return false
}

func (p *EmptyChildrenContainer) String() string {
	return "No children"
}
