package akka

import (
	"github.com/go-akka/configuration"
)

type Mailbox interface {
	Runnable

	SetActor(actor ActorCell)

	SystemEnqueue(receiver ActorRef, message SystemMessage) error
	Enqueue(receiver ActorRef, message Envelope) error

	NumberOfMessages() int
	HasMessages() bool
	HasSystemMessages() bool

	IsClosed() bool

	CanBeScheduledForExecution(hasMessageHint bool, hasSystemMessageHint bool) bool
	SetAsScheduled() bool
	SetAsIdle() bool
}

type MailboxType interface {
	Init(settings *Settings, config *configuration.Config) error
	Create(owner ActorRef, system ActorSystem) MessageQueue
}
