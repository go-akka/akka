package actor

import (
	"time"

	"github.com/go-akka/akka"
)

var (
	_ akka.Cell                = (*ActorCell)(nil)
	_ akka.UntypedActorContext = (*ActorCell)(nil)
)

type ActorCell struct {
	self       akka.InternalActorRef
	props      akka.Props
	system     akka.ActorSystem
	parent     akka.InternalActorRef
	dispitcher akka.MessageDispatcher

	currentMsg interface{}
	mailbox    akka.Mailbox

	actor *ActorBase
	IChildren
	IDispatch
}

func NewActorCell(
	system akka.ActorSystem,
	self akka.InternalActorRef,
	props akka.Props,
	dispatcher akka.MessageDispatcher,
	parent akka.InternalActorRef,
	children IChildren,
	dispatch IDispatch,
) *ActorCell {

	return &ActorCell{
		self:       self,
		system:     system,
		props:      props,
		dispitcher: dispatcher,
		parent:     parent,
		IChildren:  children,
		IDispatch:  dispatch,
	}
}

func (p *ActorCell) Dispatcher() akka.MessageDispatcher {
	return p.dispitcher
}

func (p *ActorCell) Parent() akka.ActorRef {
	return p.parent
}

func (p *ActorCell) CurrentMessage() interface{} {
	return p.currentMsg
}

func (p *ActorCell) Mailbox() akka.Mailbox {
	return p.mailbox
}

func (p *ActorCell) Self() akka.ActorRef {
	return p.self
}

func (p *ActorCell) Sender() (sender akka.ActorRef) {
	return
}

func (p *ActorCell) System() akka.ActorSystem {
	return p.system
}

func (p *ActorCell) Start() {
	return
}

func (p *ActorCell) Suspend() {
	return
}

func (p *ActorCell) Resume(err error) {
	return
}

func (p *ActorCell) Restart(err error) {
	return
}

func (p *ActorCell) StopActor(actor akka.ActorRef) (err error) {
	return
}

func (p *ActorCell) Stop() (err error) {
	return
}

func (p *ActorCell) IsLocal() bool {
	return true
}

func (p *ActorCell) Props() akka.Props {
	return p.props
}

func (p *ActorCell) HasMessages() bool {
	return false
}

func (p *ActorCell) NumberOfMessages() int {
	return 0
}

func (p *ActorCell) SendMessage(msg akka.Envelope) (err error) {
	return
}

func (p *ActorCell) IsTerminated() bool {
	return false
}

func (p *ActorCell) ChildrenRefs() akka.ChildrenContainer {
	return nil
}

func (p *ActorCell) GetSingleChild(name string) akka.ActorRef {
	return nil
}

func (p *ActorCell) GetChildByName(name string) (stats akka.ChildStats, exist bool) {
	return
}

func (p *ActorCell) UntypedBecome(behavior akka.UntypedReceive, discardOld bool) {
	return
}

func (p *ActorCell) ActorSelection(path akka.ActorPath) (selection akka.ActorSelection, err error) {
	return
}

func (p *ActorCell) Become(receive akka.ReceiveFunc, discardOld bool) (err error) {
	return
}

func (p *ActorCell) Unbecome() {
	return
}

func (p *ActorCell) ReceiveTimeout() (timeout time.Duration) {
	return
}

func (p *ActorCell) SetReceiveTimeout(timeout time.Duration) {
	return
}

func (p *ActorCell) Watch(subject akka.ActorRef) (err error) {
	return
}

func (p *ActorCell) Unwatch(subject akka.ActorRef) {
	return
}
