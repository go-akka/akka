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

func (p *UntypedActor) PreStart() (err error) {
	return p.ActorBase.PreStart()
}

func (p *UntypedActor) PostStop() (err error) {
	return p.ActorBase.PostStop()
}

func (p *UntypedActor) PreRestart(err error, message interface{}) {
	p.ActorBase.PreRestart(err, message)
}

func (p *UntypedActor) PostRestart(err error) {
	p.ActorBase.PostRestart(err)
}
