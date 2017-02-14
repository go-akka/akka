package remote

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/event"
)

var (
	_ akka.RemoteTransport = (*Remoting)(nil)
)

type EndpointPolicy interface {
}

type Remoting struct {
	system   akka.ExtendedActorSystem
	provider akka.RemoteActorRefProvider

	eventPublisher *EventPublisher
	log            akka.LoggingAdapter

	addresses      []akka.Address
	defaultAddress akka.Address
}

func NewRemoting(system akka.ExtendedActorSystem, provider akka.RemoteActorRefProvider) *Remoting {

	log := event.Logging.GetLoggerWithActorSystem(system, "remoting")
	remoting := &Remoting{
		system:         system,
		provider:       provider,
		log:            log,
		eventPublisher: NewEventPlublisher(system, log, akka.LogLevelFor(provider.RemoteSettings().RemoteLifecycleEventsLogLevel)),
	}

	return remoting
}

func (p *Remoting) Provider() akka.RemoteActorRefProvider {
	return p.provider
}

func (p *Remoting) System() akka.ExtendedActorSystem {
	return p.system
}

func (p *Remoting) Start() {
	return
}

func (p *Remoting) Addresses() []akka.Address {
	return p.addresses
}

func (p *Remoting) Send(message interface{}, sender akka.ActorRef, recipient akka.RemoteActorRef) {
	return
}

func (p *Remoting) ManagementCommand(cmd interface{}) {
	return
}

func (p *Remoting) LocalAddressForRemote(remote akka.Address) akka.Address {
	return p.defaultAddress
}

func (p *Remoting) DefaultAddress() akka.Address {
	return p.defaultAddress
}

func (p *Remoting) Log() akka.LoggingAdapter {
	return p.log
}

func (p *Remoting) Quarantine(address akka.Address, uid int, reason string) {
	return
}
