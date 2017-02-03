package actor

import (
	"github.com/go-akka/akka"
)

type UntypedActor struct {
	*ActorBase
	actor  akka.Actor
	initFn akka.InitFunc
}

func NewUntypedActor(actor akka.Actor, initFn akka.InitFunc) *UntypedActor {

	untypedActor := &UntypedActor{
		actor:  actor,
		initFn: initFn,
	}

	return untypedActor
}

func (p *UntypedActor) Construct() error {
	if p.initFn != nil {
		return p.initFn()
	}
	return nil
}

func (p *UntypedActor) SetActorBase(actorBase *ActorBase) {
	p.ActorBase = actorBase
}

func (p *UntypedActor) Receive(message interface{}) (handled bool, err error) {
	return p.actor.Receive(message)
}
