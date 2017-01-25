package actor

import (
	"github.com/go-akka/akka"
	"time"
)

var emptyBehavior = func(_ interface{}) bool {
	return false
}

type ActorBase struct {
	clearedSelf    akka.ActorRef
	hasBeenCleared bool

	actor        akka.Actor
	receiverFunc akka.ReceiveFunc

	ctx akka.ActorContext
}

func NewActorBase(receiverFunc akka.ReceiveFunc, actor akka.Actor) *ActorBase {
	actorBase := &ActorBase{
		actor:        actor,
		receiverFunc: receiverFunc,
	}

	return actorBase
}

func (p *ActorBase) String() string {
	return ""
}

func (p *ActorBase) Context() akka.ActorContext {
	if p.hasBeenCleared {
		return nil
	}
	return p.ctx
}

func (p *ActorBase) AroundReceive(receiveFunc akka.ReceiveFunc, message interface{}) (wasHandled bool, err error) {
	wasHandled, err = receiveFunc(message)
	if !wasHandled {
		p.Unhandled(message)
	}
	return
}

func (p *ActorBase) Receive(message interface{}) (wasHandled bool, err error) {

	if wasHandled, err = p.receiverFunc(message); err != nil {
		return
	} else if !wasHandled {
		err = p.Unhandled(message)
		return
	}

	return
}

func (p *ActorBase) Unhandled(message interface{}) (err error) {
	if terminatedMessage, ok := message.(akka.Terminated); ok {
		p.Context().System().EventStream().Publish(&akka.UnhandledMessage{terminatedMessage, p.Sender(), p.Self()})
	}

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
