package actor

import (
	"github.com/go-akka/akka"
)

var (
	_ akka.LocalActorRef = (*LocalActorRef)(nil)
)

type LocalActorRef struct {
	akka.LocalRef

	system      *ActorSystemImpl
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

	sysImpl := system.(*ActorSystemImpl)

	ref := &LocalActorRef{
		system:      sysImpl,
		props:       props,
		dispatcher:  dispatcher,
		mailboxType: mailboxType,
		supervisor:  supervisor,
		path:        path,
	}

	actorCell := newActorCell(sysImpl, ref, props, dispatcher, supervisor, nil, nil)

	actorCell.Init(true, mailboxType)

	ref.cell = actorCell

	return ref
}

func (p *LocalActorRef) Tell(message interface{}, sender ...akka.ActorRef) error {
	var s akka.ActorRef
	if len(sender) == 0 {
		s = (*akka.NoSender)(nil)
	} else {
		s = sender[0]
	}

	return p.cell.SendMessage(akka.Envelope{message, s})
}

func (p *LocalActorRef) Path() akka.ActorPath {
	return nil
}

func (p *LocalActorRef) CompareTo(other akka.ActorRef) int {
	return 0
}

func (p *LocalActorRef) String() string {
	return "LocalActorRef"
}

func (p *LocalActorRef) Parent() akka.InternalActorRef {
	return nil
}

func (p *LocalActorRef) GetChild(names ...string) akka.InternalActorRef {
	return nil
}

func (p *LocalActorRef) Resume(causedByFailure error) {
	p.cell.Resume(causedByFailure)
}

func (p *LocalActorRef) Start() {
	p.cell.Start()
}

func (p *LocalActorRef) Stop() {
	p.cell.Stop()
}

func (p *LocalActorRef) Suspend() {
	p.cell.Suspend()
}

func (p *LocalActorRef) IsTerminated() bool {
	return p.cell.IsTerminated()
}

func (p *LocalActorRef) Underlying() akka.Cell {
	return p.cell
}

func (p *LocalActorRef) Cell() *ActorCell {
	return p.cell
}

func (p *LocalActorRef) Children() []akka.ActorRef {
	return p.cell.Children()
}

func (p *LocalActorRef) GetSingleChild(name string) akka.InternalActorRef {
	ref := p.cell.GetSingleChild(name)
	internalActorRef, _ := ref.(akka.InternalActorRef)
	return internalActorRef
}

func (p *LocalActorRef) GetParent() akka.InternalActorRef {
	return p.cell.parent
}

func (p *LocalActorRef) Provider() akka.ActorRefProvider {
	return p.cell.Provider()
}

func (p *LocalActorRef) SendSystemMessage(message akka.SystemMessage) (err error) {
	return p.cell.SendSystemMessage(message)
}

func (p *LocalActorRef) Restart(cause error) {
	p.cell.Restart(cause)
}
