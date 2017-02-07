package actor

import (
	"context"
	"fmt"
	"time"

	"github.com/go-akka/akka"
)

const (
	contextKeyOfActorContext = "akka.ActorContext"
)

var emptyBehavior = func(_ interface{}) bool {
	return false
}

type actorBaseSetter interface {
	SetActorBase(actorBase *ActorBase)
}

type ActorBase struct {
	clearedSelf    akka.ActorRef
	hasBeenCleared bool

	actor akka.Actor

	context context.Context
}

func NewActorBase(actor akka.Actor, ctx akka.ActorContext) *ActorBase {
	actorBase := &ActorBase{
		actor:   actor,
		context: context.WithValue(context.Background(), contextKeyOfActorContext, ctx),
	}

	return actorBase
}

func (p *ActorBase) Context() akka.ActorContext {
	if p.hasBeenCleared {
		return nil
	}

	v := p.context.Value(contextKeyOfActorContext)

	if ctx, ok := v.(akka.ActorContext); ok {
		return ctx
	}

	return nil
}

func (p *ActorBase) AroundReceive(receiveFunc akka.ReceiveFunc, message interface{}) (wasHandled bool, err error) {
	wasHandled, err = receiveFunc(message)
	if !wasHandled {
		p.Unhandled(message)
	}
	return
}

func (p *ActorBase) Receive(message interface{}) (wasHandled bool, err error) {

	if wasHandled, err = p.actor.Receive(message); err != nil {
		return
	} else if !wasHandled {
		err = p.Unhandled(message)
		return
	}

	return
}

func (p *ActorBase) Unhandled(message interface{}) (err error) {
	if terminatedMessage, ok := message.(*Terminated); ok {
		err = fmt.Errorf("Monitored actor [%s] terminated", terminatedMessage.Actor)
		return
	}

	p.Context().System().EventStream().Publish(&akka.UnhandledMessage{message, p.Sender(), p.Self()})

	return
}

func (p *ActorBase) Sender() akka.ActorRef {
	return p.Context().Sender()
}

func (p *ActorBase) Self() akka.ActorRef {
	if p.hasBeenCleared {
		return p.clearedSelf
	}
	return p.Context().Self()
}

func (p *ActorBase) Become(receive akka.ReceiveFunc, discardOld bool) (err error) {
	return p.Context().Become(receive, discardOld)
}

func (p *ActorBase) SetReceiveTimeout(timeout time.Duration) {
	p.Context().SetReceiveTimeout(timeout)
}
