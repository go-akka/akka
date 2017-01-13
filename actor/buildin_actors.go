package actor

import (
	"github.com/go-akka/akka"
)

type EventStreamActor struct {
	*ActorBase
}

func (p *EventStreamActor) Receive(message interface{}) (unhandled bool) {
	return
}

type GuardianActor struct {
	*ActorBase
}

func (p *GuardianActor) Receive(message interface{}) (unhandled bool) {

	switch msg := message.(type) {
	case akka.Terminated:
		{
			p.Context().StopActor(p.Self())
			return
		}
	case akka.StopChild:
		{
			p.Context().StopActor(msg.Child())
		}
	default:
		p.Context().System().DeadLetters().Tell(akka.NewDeadLetter(message, p.Sender(), p.Self()), p.Sender())
	}

	return
}

func (p *GuardianActor) PreStart() (err error) {
	return
}

type SystemGuardianActor struct {
	*ActorBase

	userGuardian     akka.ActorRef
	terminationHooks map[akka.ActorRef]bool
}

func (p *SystemGuardianActor) Receive(message interface{}) (unhandled bool) {

	switch msg := message.(type) {
	case *akka.Terminated:
		{
			terminatedActor := msg.Actor()
			if p.userGuardian.Equals(terminatedActor) {
				p.Context().Become(p.Terminating, true)

				for terminationHook, _ := range p.terminationHooks {
					terminationHook.Tell(akka.TerminationHook{}, akka.NoSender{})
				}

				p.stopWhenAllTerminationHooksDone()
			} else {
				delete(p.terminationHooks, terminatedActor)
			}

			return
		}
	case *akka.StopChild:
		{
			p.Context().StopActor(msg.Child())
		}
	default:
		p.Context().System().DeadLetters().Tell(akka.NewDeadLetter(message, p.Sender(), p.Self()), p.Sender())
	}

	return
}

func (p *SystemGuardianActor) Terminating(message interface{}) (unhandled bool, err error) {
	return
}

func (p *SystemGuardianActor) stopWhenAllTerminationHooksDone() {
	if len(p.terminationHooks) == 0 {
		p.Context().StopActor(p.Self())
	}
}

func (p *SystemGuardianActor) PreStart() (err error) {
	return
}
