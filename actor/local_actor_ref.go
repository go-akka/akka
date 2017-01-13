package actor

import (
	"github.com/go-akka/akka"
)

var (
	_ akka.LocalActorRef = (*LocalActorRef)(nil)
)

type LocalActorRef struct {
	akka.LocalRef

	system      akka.ActorSystem
	props       akka.Props
	dispatcher  akka.MessageDispatcher
	mailboxType akka.MailboxType
	supervisor  akka.InternalActorRef
	path        akka.ActorPath
	cell        *ActorCell
}

func NewLocalActorRef(
	system akka.ActorSystem,
	props akka.Props,
	dispatcher akka.MessageDispatcher,
	mailboxType akka.MailboxType,
	supervisor akka.InternalActorRef,
	path akka.ActorPath,
) *LocalActorRef {
	return &LocalActorRef{
		system:      system,
		props:       props,
		dispatcher:  dispatcher,
		mailboxType: mailboxType,
		supervisor:  supervisor,
		path:        path,
	}
}

func (p *LocalActorRef) Tell(message interface{}, sender akka.ActorRef) {
	return
}

func (p *LocalActorRef) Path() akka.ActorPath {
	return nil
}

func (p *LocalActorRef) Equals(that interface{}) bool {
	return false
}

func (p *LocalActorRef) Provider() akka.ActorRefProvider {
	return nil
}

func (p *LocalActorRef) String() string {
	return "LocalActorRef"
}

func (p *LocalActorRef) Parent() akka.InternalActorRef {
	return nil
}

func (p *LocalActorRef) Child(names ...string) akka.InternalActorRef {
	return nil
}

func (p *LocalActorRef) Resume(err error) {

}

func (p *LocalActorRef) Start() {

}

func (p *LocalActorRef) Stop() {

}

func (p *LocalActorRef) Restart(err error) {
	return
}

func (p *LocalActorRef) Suspend() {
}

func (p *LocalActorRef) Underlying() akka.Cell {
	return p.cell
}

func (p *LocalActorRef) Cell() *ActorCell {
	return p.cell
}

func (p *LocalActorRef) Children() []akka.ActorRef {
	return nil
}

func (p *LocalActorRef) GetSingleChild(name string) akka.InternalActorRef {
	return nil
}
