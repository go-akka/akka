package internal

import (
	"github.com/go-akka/akka"
)

type ChildrenContainerBase struct {
	children *ImmutableMap
}

func (p *ChildrenContainerBase) Children() []akka.ActorRef {
	var refs []akka.ActorRef
	for _, v := range p.children.Items() {
		if item, ok := v.(akka.ChildRestartStats); ok {
			refs = append(refs, item.Child())
		}
	}
	return refs
}

func (p *ChildrenContainerBase) Stats() []akka.ChildRestartStats {
	var children []akka.ChildRestartStats
	for _, v := range p.children.Items() {
		if item, ok := v.(akka.ChildRestartStats); ok {
			children = append(children, item)
		}
	}

	return children
}

func (p *ChildrenContainerBase) GetByName(name string) (stats akka.ChildStats, exist bool) {
	if item, exist := p.children.Get(name); exist {
		return item.(akka.ChildStats), true
	}
	return
}

func (p *ChildrenContainerBase) GetByRef(actor akka.ActorRef) (crStats akka.ChildRestartStats, exist bool) {
	if v, ok := p.children.Get(actor.Path().Name()); ok {
		if stats, ok := v.(akka.ChildRestartStats); ok {
			if actor.CompareTo(stats.Child()) == 0 {
				crStats = stats
				exist = true
				return
			}
		}
	}

	return
}

func (p *ChildrenContainerBase) Contains(actor akka.ActorRef) bool {
	_, exist := p.GetByRef(actor)
	return exist
}

func (p *ChildrenContainerBase) IsTerminating() bool {
	return false
}

func (p *ChildrenContainerBase) IsNormal() bool {
	return true
}
