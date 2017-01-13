package actor

import (
	"github.com/go-akka/akka"
)

type IDispatch interface {
	Init(sendSupervise bool, mailboxType akka.MailboxType)
	InitWithFailure(err error)
	Mailbox() (mailbox akka.Mailbox)
	HasMessages() (has bool)
	NumberOfMessages() int
	IsTerminated() (yes bool)
	Start()
}

var (
	_ IDispatch = (*ActorCellDispatch)(nil)
)

type ActorCellDispatch struct {
	cell *ActorCell
}

func NewActorCellDispatch(cell *ActorCell) IDispatch {
	return &ActorCellDispatch{cell: cell}
}

func (p *ActorCellDispatch) Init(sendSupervise bool, mailboxType akka.MailboxType) {
	return
}

func (p *ActorCellDispatch) InitWithFailure(err error) {
	return
}

func (p *ActorCellDispatch) Mailbox() (mailbox akka.Mailbox) {
	return
}

func (p *ActorCellDispatch) HasMessages() (has bool) {
	return
}

func (p *ActorCellDispatch) NumberOfMessages() int {
	return 0
}

func (p *ActorCellDispatch) IsTerminated() (yes bool) {
	return
}

func (p *ActorCellDispatch) Start() {
	return
}
