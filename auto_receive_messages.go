package akka

import (
	"fmt"
)

type Kill struct{}
type PoisonPill struct{}
type AddressTerminated struct{ Address Address }
type ActorSelectionMessage struct{}
type Identify struct{ MessageID interface{} }

func (p *PoisonPill) AutoReceivedMessage() {}
func (p *PoisonPill) String() string {
	return "<PoisonPill>"
}

func (p *Kill) AutoReceivedMessage() {}
func (p *Kill) String() string {
	return "<Kill>"
}

func (p *AddressTerminated) AutoReceivedMessage() {}
func (p *AddressTerminated) String() string {
	return "<AddressTerminated>:" + p.Address.String()
}

func (p *ActorSelectionMessage) AutoReceivedMessage() {}
func (p *ActorSelectionMessage) String() string {
	return "<ActorSelectionMessage>"
}

func (p *Identify) AutoReceivedMessage() {}
func (p *Identify) String() string {
	return fmt.Sprintf("<Identify>: %v", p.MessageID)
}

type Terminated struct {
	Actor              ActorRef
	AddressTerminated  bool
	ExistenceConfirmed bool
}

func (p *Terminated) AutoReceivedMessage() {}
func (p *Terminated) String() string {
	strExistenceConfirmed := "False"
	if p.ExistenceConfirmed {
		strExistenceConfirmed = "True"
	}
	return "<Terminated>: " + p.Actor.Path().String() + " - ExistenceConfirmed=" + strExistenceConfirmed
}
