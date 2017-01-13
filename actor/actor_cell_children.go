package actor

import (
	"github.com/go-akka/akka"
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
	cell *ActorCell
}

func NewActorCellChildren(cell *ActorCell) IChildren {
	return &ActorCellChildren{cell: cell}
}

func (p *ActorCellChildren) Children() []akka.ActorRef {
	return nil
}

func (p *ActorCellChildren) Child(name string) (ref akka.ActorRef, exist bool) {
	return
}

func (p *ActorCellChildren) ActorOf(props akka.Props, name string) (ref akka.ActorRef, err error) {
	return
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
	return
}
