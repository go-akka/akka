package actor

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/actor/internal"
	"math/rand"
	"sync"
)

type IChildren interface {
	Children() []akka.ActorRef
	Child(name string) (ref akka.ActorRef, exist bool)

	ActorOf(props akka.Props, name string) (ref akka.ActorRef, err error)
	StopChild(actor akka.ActorRef)

	ReserveChild(name string) bool
	InitChild(ref akka.ActorRef) *akka.ChildRestartStats

	AttachChild(props akka.Props, name string, systemService bool) (akka.ActorRef, error)
	ChildrenRefs() akka.ChildrenContainer
}

var (
	_ IChildren = (*ActorCellChildren)(nil)
)

type ActorCellChildren struct {
	*ActorCell

	childrenContainer akka.ChildrenContainer

	containerLocker sync.Mutex
}

func newActorCellChildren(cell *ActorCell) IChildren {
	return &ActorCellChildren{ActorCell: cell, childrenContainer: internal.EmptyChildrenContainerInstance}
}

func (p *ActorCellChildren) Children() []akka.ActorRef {
	return p.childrenContainer.Children()
}

func (p *ActorCellChildren) ChildrenRefs() akka.ChildrenContainer {
	return p.childrenContainer
}

func (p *ActorCellChildren) Child(name string) (ref akka.ActorRef, exist bool) {
	return
}

func (p *ActorCellChildren) ActorOf(props akka.Props, name string) (ref akka.ActorRef, err error) {
	// TODO: check name
	return p.makeChild(props, name, false, false)
}

func (p *ActorCellChildren) StopChild(actor akka.ActorRef) {
	_, exist := p.childrenContainer.GetByRef(actor)
	if exist {
		// TODO
	}

	(actor.(akka.InternalActorRef)).Stop()
}

func (p *ActorCellChildren) ReserveChild(name string) bool {
	return p.updateChildrenRefs(p.childrenContainer.Reserve(name))
}

func (p *ActorCellChildren) UnreserveChild(name string) bool {
	return p.updateChildrenRefs(p.childrenContainer.Unreserve(name))
}

func (p *ActorCellChildren) InitChild(ref akka.ActorRef) *akka.ChildRestartStats {
	return nil
}

func (p *ActorCellChildren) AttachChild(props akka.Props, name string, systemService bool) (ref akka.ActorRef, err error) {
	return p.makeChild(props, name, true, systemService)
}

func (p *ActorCellChildren) NewUID() int64 {
	uid := rand.Int63()
	for uid == 0 {
		uid = rand.Int63()
	}
	return uid
}

func (p *ActorCellChildren) makeChild(props akka.Props, name string, async bool, systemService bool) (ref akka.ActorRef, err error) {

	p.ReserveChild(name)
	var actor akka.InternalActorRef

	childPath := akka.NewChildActorPath(p.Self().Path(), name, p.NewUID())

	actor = p.system.provider.ActorOf(p.system, props, p.self, childPath, systemService, nil, true, async)

	// if p.Mailbox() != nil {

	// }

	p.InitChild(actor)
	actor.Start()

	ref = actor

	return
}

func (p *ActorCellChildren) updateChildrenRefs(newRef akka.ChildrenContainer) bool {
	p.containerLocker.Lock()
	defer p.containerLocker.Unlock()

	if p.childrenContainer == newRef {
		return false
	}

	p.childrenContainer = newRef

	return true
}
