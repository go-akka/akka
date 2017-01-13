package actor

import (
	"github.com/go-akka/akka"
)

type UntypedActor struct {
	*ActorBase

	receive akka.ReceiveFunc
}

func NewUntypedActor(receive akka.ReceiveFunc) *UntypedActor {

	untypedActor := &UntypedActor{
		receive: receive,
	}

	untypedActor.ActorBase = NewActorBase(untypedActor.Receive)

	return untypedActor
}

func (p *UntypedActor) Receive(message interface{}) (handled bool, err error) {
	return p.receive(message)
}
