package actor

import (
	"github.com/go-akka/akka"
)

func (p *ActorBase) AroundPreReStart(cause error, message interface{}) {
	p.PreRestart(cause, message)
}

func (p *ActorBase) AroundPreStart() (err error) {
	if preStarter, ok := p.actor.(akka.PreStarter); ok {
		return preStarter.PreStart()
	}
	return
}

func (p *ActorBase) AroundPostRestart(cause error, message interface{}) {
	p.PostRestart(cause)
}

func (p *ActorBase) PreRestart(cause error, message interface{}) {
	for _, child := range p.Context().Children() {
		p.Context().Unwatch(child)
		p.Context().StopActor(child)
	}
	p.PostStop()
}

func (p *ActorBase) PostRestart(cause error) {
	// p.PreStart()
}

func (p *ActorBase) AroundPostStop() {
	p.PostStop()
}

func (p *ActorBase) PostStop() (err error) {
	return
}
