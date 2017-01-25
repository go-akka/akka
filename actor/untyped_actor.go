package actor

import (
	"github.com/go-akka/akka"
)

type UntypedActor struct {
	*ActorBase
	actor akka.Actor
}

func NewUntypedActor(actor akka.Actor) *UntypedActor {

	untypedActor := &UntypedActor{
		actor: actor,
	}

	untypedActor.ActorBase = NewActorBase(untypedActor.Receive, actor)

	return untypedActor
}

func (p *UntypedActor) Receive(message interface{}) (handled bool, err error) {
	return p.actor.Receive(message)
}
