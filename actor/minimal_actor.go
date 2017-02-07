package actor

import (
	"github.com/go-akka/akka"
)

type MinimalActor struct {
	*ActorBase
	receiver akka.ContextReceiver
	initFn   akka.InitFunc
}

func NewMinimalActor(receiver akka.ContextReceiver, initFn akka.InitFunc) *MinimalActor {

	minimalActor := &MinimalActor{
		receiver: receiver,
		initFn:   initFn,
	}

	return minimalActor
}

func (p *MinimalActor) construct() error {
	if p.initFn != nil {
		return p.initFn()
	}
	return nil
}

func (p *MinimalActor) SetActorBase(actorBase *ActorBase) {
	p.ActorBase = actorBase
}

func (p *MinimalActor) PreStart() (err error) {
	switch preStarter := p.receiver.(type) {
	case akka.PreStarter:
		{
			return preStarter.PreStart()
		}
	case akka.ContextPreStarter:
		{
			return preStarter.PreStart(p.Context())
		}
	}
	return
}

func (p *MinimalActor) Receive(message interface{}) (handled bool, err error) {
	return p.receiver.Receive(p.Context(), message)
}
