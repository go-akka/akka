package actor

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/dispatch/sysmsg"
)

type RootGuardianActor struct {
	*UntypedActor
}

func (p *RootGuardianActor) Construct(supervisorStrategy akka.SupervisorStrategy) {

}

func (p *RootGuardianActor) Receive(message interface{}) (handled bool, err error) {
	switch message.(type) {
	case *Terminated:
	case *sysmsg.StopChild:
	}

	return
}

func (p *RootGuardianActor) PreRestart(cause error, message interface{}) {
	return
}

func (p *RootGuardianActor) PreStart() (err error) {
	return
}
