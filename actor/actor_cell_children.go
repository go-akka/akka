package actor

import (
	"github.com/go-akka/akka"
	"math/rand"
)

type IChildren interface {
	Children() []akka.ActorRef
	Child(name string) (ref akka.ActorRef, exist bool)

	ActorOf(props akka.Props, name string) (ref akka.ActorRef, err error)
	StopChild(actor akka.ActorRef)

	ReserveChild(name string) bool
	InitChild(ref akka.ActorRef) *akka.ChildRestartStats

	AttachChild(props akka.Props, name string, systemService bool) (akka.ActorRef, error)
}

var (
	_ IChildren = (*ActorCellChildren)(nil)
)

type ActorCellChildren struct {
	*ActorCell
}

func newActorCellChildren(cell *ActorCell) IChildren {
	return &ActorCellChildren{cell}
}

func (p *ActorCellChildren) Children() []akka.ActorRef {
	return nil
}

func (p *ActorCellChildren) Child(name string) (ref akka.ActorRef, exist bool) {
	return
}

func (p *ActorCellChildren) ActorOf(props akka.Props, name string) (ref akka.ActorRef, err error) {
	return p.makeChild(props, name, false, false)
}

func (p *ActorCellChildren) StopChild(actor akka.ActorRef) {
	return
}

func (p *ActorCellChildren) ReserveChild(name string) bool {
	return false
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
