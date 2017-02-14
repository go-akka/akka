package remote

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/actor"
	"time"
)

type Quarantine struct {
	Uid           int
	RemoteAddress akka.Address
}

type Quarantined struct {
	Uid      int
	Deadline time.Time
}

type ListensResult struct {
}

type ListensFailure struct {
}

type EndpointManager struct {
	*actor.UntypedActor

	settings         *akka.RemoteSettings
	extendedSystem   akka.ExtendedActorSystem
	endpointId       []int
	eventPublisher   *EventPublisher
	endpoints        *EndpointRegistry
	transportMapping map[akka.Address]AkkaProtocolTransport
}

func (p *EndpointManager) Receive(message interface{}) (handled bool, err error) {
	switch message.(type) {
	case *Listen:
		{

		}
	case *ListensResult:
		{

		}
	case *ListensFailure:
		{

		}
	case *InboundAssociation:
		{

		}
	case *ManagementCommand:
		{
			p.Sender().Tell(&ManagementCommandAck{Status: true})
		}
	case *StartupFinished:
		{
			p.Become(p.AcceptingReceive, true)
		}
	case *ShutdownAndFlush:
		{
			p.Sender().Tell(true)
			p.Context().StopChild(p.Self())
		}
	}

	return
}

func (p *EndpointManager) AcceptingReceive(message interface{}) (handled bool, err error) {
	switch message.(type) {
	case *ManagementCommand:
		{

		}
	case *Quarantine:
		{

		}
	case *Send:
		{

		}
	case *ShutdownAndFlush:
		{
			p.Become(p.FlushingReceive, true)
		}
	}

	return
}

func (p *EndpointManager) FlushingReceive(message interface{}) (handled bool, err error) {
	switch message.(type) {
	case *Send:
		{
		}
	case *InboundAssociation:
		{
		}
	case *actor.Terminated:
		{
		}
	}

	return
}
