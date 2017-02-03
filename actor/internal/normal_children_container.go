package internal

import (
	"github.com/go-akka/akka"
)

type NormalChildrenContainer struct {
	*ChildrenContainerBase
}

func createNormalChildContainer(children *ImmutableMap) akka.ChildrenContainer {
	if children.Count() == 0 {
		return EmptyChildrenContainerInstance
	}

	return newNormalChildrenContainer(children.copy())

}

func newNormalChildrenContainer(children *ImmutableMap) akka.ChildrenContainer {
	return &NormalChildrenContainer{
		ChildrenContainerBase: &ChildrenContainerBase{children: children},
	}
}

func (p *NormalChildrenContainer) Add(name string, stats akka.ChildRestartStats) akka.ChildrenContainer {
	return createNormalChildContainer(p.children.Set(name, stats))
}

func (p *NormalChildrenContainer) Remove(child akka.ActorRef) akka.ChildrenContainer {
	return createNormalChildContainer(p.children.Remove(child.Path().Name()))
}

func (p *NormalChildrenContainer) ShallDie(actor akka.ActorRef) akka.ChildrenContainer {
	return p
}

func (p *NormalChildrenContainer) Reserve(name string) akka.ChildrenContainer {
	// TODO: throw error while key exist
	return newNormalChildrenContainer(p.children.Set(name, _childNameReservedInstance))
}

func (p *NormalChildrenContainer) Unreserve(name string) akka.ChildrenContainer {
	item, exist := p.children.Get(name)
	if !exist {
		return p
	}

	if _, ok := item.(akka.ChildNameReserved); ok {
		return createNormalChildContainer(p.children.Remove(name))
	}

	return p
}

func (p *NormalChildrenContainer) String() string {
	return "NormalChildrenContainer"
}
