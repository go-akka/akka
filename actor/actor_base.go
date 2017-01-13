package actor

import (
	"github.com/go-akka/akka"
	"sync"
)

var emptyBehavior = func(_ interface{}) bool {
	return false
}

type ActorBase struct {
	clearedSelf    akka.ActorRef
	hasBeenCleared bool

	receive akka.ReceiveFunc

	cellInitOnce sync.Once

	ctx akka.ActorContext
}

func NewActorBase(receive akka.ReceiveFunc) *ActorBase {
	actorBase := &ActorBase{
		receive: receive,
	}

	actorBase.Become(receive, true)

	return actorBase
}

func (p *ActorBase) String() string {
	return ""
}

func (p *ActorBase) Context() akka.ActorContext {
	p.cellInitOnce.Do(func() {
		p.ctx = &ActorCell{}
	})

	return p.ctx
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

func (p *ActorBase) Receive(message interface{}) (wasHandled bool, err error) {

	if wasHandled, err = p.receive(message); err != nil {
		return
	} else if !wasHandled {
		err = p.Unhandled(message)
		return
	}

	return
}

func (p *ActorBase) Unhandled(message interface{}) (err error) {
	return
}
