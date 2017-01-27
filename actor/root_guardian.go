package actor

import (
	"github.com/go-akka/akka"
)

type RootGuardianActor struct {
	*UntypedActor
}

func (p *RootGuardianActor) Construct(supervisorStrategy akka.SupervisorStrategy) {

}

func (p *RootGuardianActor) Receive(message interface{}) (handled bool, err error) {
	switch message.(type) {
	case *akka.Terminated:
	case *akka.StopChild:
	}

	return
}

func (p *RootGuardianActor) PreRestart(cause error, message interface{}) {
	return
}

func (p *RootGuardianActor) PreStart() (err error) {
	return
}
