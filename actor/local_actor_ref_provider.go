package actor

import (
	"github.com/go-akka/akka"
)

type LocalActorRefProvider struct {
	settings   akka.Settings
	eventStrem akka.EventStream
	deployer   akka.Deployer
}

func NewLocalActorRefProvider(systemName string,
	settings akka.Settings,
	eventStrem akka.EventStream,
	deployer akka.Deployer,
) akka.ActorRefProvider {
	return &LocalActorRefProvider{
		settings:   settings,
		eventStrem: eventStrem,
		deployer:   deployer,
	}
}

func (p *LocalActorRefProvider) ActorOf(
	system akka.ActorSystem,
	props akka.Props,
	supervisor akka.InternalActorRef,
	path akka.ActorPath,
	systemService bool,
	deploy akka.Deploy,
	lookupDeploy bool,
	async bool) akka.InternalActorRef {
	// dispitcher := akka.MessageDispatcher{}
	return NewLocalActorRef(system, props, nil, nil, supervisor, path)
}

func (p *LocalActorRefProvider) Init(system akka.ActorSystem) {
	return
}

func (p *LocalActorRefProvider) DeadLetters() akka.ActorRef {
	return nil
}

func (p *LocalActorRefProvider) Deployer() akka.Deployer {
	return p.deployer
}

func (p *LocalActorRefProvider) DefaultAddress() akka.Address {
	return akka.Address{}
}

func (p *LocalActorRefProvider) ExternalAddressFor(addr akka.Address) akka.Address {
	return akka.Address{}
}

func (p *LocalActorRefProvider) Guardian() akka.LocalActorRef {
	return nil
}

func (p *LocalActorRefProvider) RegisterTempActor(actorRef akka.InternalActorRef, path akka.ActorPath) {
	return
}

func (p *LocalActorRefProvider) ResolveActorRef(path akka.ActorPath) akka.ActorRef {
	return nil
}

func (p *LocalActorRefProvider) RootGuardian() akka.InternalActorRef {
	return nil
}

func (p *LocalActorRefProvider) RootGuardianAt(address akka.Address) akka.ActorRef {
	return nil
}

func (p *LocalActorRefProvider) RootPath() akka.ActorPath {
	return nil
}

func (p *LocalActorRefProvider) Settings() akka.Settings {
	return p.settings
}

func (p *LocalActorRefProvider) SystemGuardian() akka.LocalActorRef {
	return nil
}

func (p *LocalActorRefProvider) TempContainer() akka.InternalActorRef {
	return nil
}

func (p *LocalActorRefProvider) TempPath() akka.ActorPath {
	return nil
}

func (p *LocalActorRefProvider) TerminationFuture() {
	return
}

func (p *LocalActorRefProvider) UnregisterTempActor(path akka.ActorPath) {
	return
}
