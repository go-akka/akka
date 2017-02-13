package actor

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/actor/props"
	"github.com/go-akka/akka/dispatch"
	"github.com/go-akka/akka/pkg/dynamic_access"
	"sync"
)

var (
	_ akka.LocalActorRefProvider = (*LocalActorRefProvider)(nil)
)

type LocalActorRefProvider struct {
	systemName    string
	system        *ActorSystemImpl
	settings      *akka.Settings
	eventStrem    akka.EventStream
	deployer      akka.Deployer
	dynamicAccess dynamic_access.DynamicAccess

	defaultDispatcher akka.MessageDispatcher
	defaultMailbox    akka.MailboxType

	rootPath       akka.ActorPath
	rootGuardian   akka.LocalActorRef
	guardian       akka.LocalActorRef
	systemGuardian akka.LocalActorRef

	constructOnce sync.Once
}

func (p *LocalActorRefProvider) Construct(
	systemName string,
	settings *akka.Settings,
	eventStrem akka.EventStream,
	dynamicAccess dynamic_access.DynamicAccess) {

	p.constructOnce.Do(func() {
		p.systemName = systemName
		p.settings = settings
		p.eventStrem = eventStrem
		p.dynamicAccess = dynamicAccess
	})
}

func (p *LocalActorRefProvider) Init(system akka.ActorSystem) (err error) {
	p.system = system.(*ActorSystemImpl)
	p.defaultDispatcher = p.system.dispatchers.Lookup(dispatch.DefaultDispatcherId)
	p.defaultMailbox, _ = p.system.mailboxes.Lookup(dispatch.DefaultMailboxId)

	p.rootPath = akka.NewRootActorPath(akka.NewAddress("akka", p.systemName, "", 0), "/")

	var rootGuardian, userGuardian, systemGuardian akka.LocalActorRef

	if rootGuardian, err = p.createRootGuardian(system); err != nil {
		return
	}

	if userGuardian, err = p.createUserGuardian(rootGuardian, "user"); err != nil {
		return
	}

	if systemGuardian, err = p.createSystemGuardian(rootGuardian, "system", userGuardian); err != nil {
		return
	}

	p.rootGuardian = rootGuardian
	p.guardian = userGuardian
	p.systemGuardian = systemGuardian

	p.rootGuardian.Start()

	p.eventStrem.StartDefaultLoggers(p.system)

	return
}

func (p *LocalActorRefProvider) ActorOf(
	system akka.ActorSystem,
	props akka.Props,
	supervisor akka.InternalActorRef,
	path akka.ActorPath,
	systemService bool,
	deploy *akka.Deploy,
	lookupDeploy bool,
	async bool) akka.InternalActorRef {

	sys := system.(*ActorSystemImpl)

	dispatcher := sys.dispatchers.Lookup(props.Dispatcher())
	mailboxType, _ := sys.mailboxes.Lookup(props.Mailbox())

	// dispitcher := akka.MessageDispatcher{}
	return NewLocalActorRef(sys, props, dispatcher, mailboxType, supervisor, path)
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
	return p.guardian
}

func (p *LocalActorRefProvider) SystemGuardian() akka.LocalActorRef {
	return p.systemGuardian
}

func (p *LocalActorRefProvider) RegisterTempActor(actorRef akka.InternalActorRef, path akka.ActorPath) {
	return
}

func (p *LocalActorRefProvider) ResolveActorRef(path akka.ActorPath) akka.ActorRef {
	return nil
}

func (p *LocalActorRefProvider) RootGuardian() akka.InternalActorRef {
	return p.rootGuardian
}

func (p *LocalActorRefProvider) RootGuardianAt(address akka.Address) akka.ActorRef {
	return nil
}

func (p *LocalActorRefProvider) RootPath() akka.ActorPath {
	return nil
}

func (p *LocalActorRefProvider) Settings() *akka.Settings {
	return p.settings
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

func (p *LocalActorRefProvider) LocalActorRefProvider() {}

func (p *LocalActorRefProvider) createRootGuardian(system akka.ActorSystem) (ref akka.LocalActorRef, err error) {

	var actorProps akka.Props
	actorProps, err = props.Create((*RootGuardianActor)(nil), nil)
	if err != nil {
		return
	}

	theOneWhoWalksTheBubblesOfSpaceTime := NewBubbleWalker(p.rootPath.Append("bubble-walker"), p)

	ref = NewLocalActorRef(system, actorProps, p.defaultDispatcher, p.defaultMailbox, theOneWhoWalksTheBubblesOfSpaceTime, p.rootPath)

	return
}

func (p *LocalActorRefProvider) createUserGuardian(rootGuardian akka.LocalActorRef, name string) (ref akka.LocalActorRef, err error) {
	cell := rootGuardian.Underlying().(*ActorCell)
	cell.ReserveChild(name)

	var actorProps akka.Props
	actorProps, err = props.Create((*RootGuardianActor)(nil), nil)
	if err != nil {
		return
	}

	userGuardian := NewLocalActorRef(p.system, actorProps, p.defaultDispatcher, p.defaultMailbox, rootGuardian, p.rootPath.Append(name))

	cell.InitChild(userGuardian)
	userGuardian.Start()

	ref = userGuardian
	return
}

func (p *LocalActorRefProvider) createSystemGuardian(rootGuardian akka.LocalActorRef, name string, userGuardian akka.LocalActorRef) (ref akka.LocalActorRef, err error) {
	cell := rootGuardian.Underlying().(*ActorCell)
	cell.ReserveChild(name)

	var actorProps akka.Props
	actorProps, err = props.Create((*SystemGuardianActor)(nil), userGuardian)
	if err != nil {
		return
	}

	systemGuardian := NewLocalActorRef(p.system, actorProps, p.defaultDispatcher, p.defaultMailbox, rootGuardian, p.rootPath.Append(name))

	cell.InitChild(systemGuardian)
	systemGuardian.Start()

	ref = systemGuardian
	return

}
