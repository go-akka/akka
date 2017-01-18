package akka

import (
	"github.com/go-akka/configuration"
)

type Mailbox interface {
	Runnable
}

// type MailboxType struct {
// 	settings Settings
// 	config   *configuration.Config
// }

type MailboxType interface {
	Init(settings Settings, config configuration.Config) error
	Create(owner ActorRef, system ActorSystem) MessageQueue
}

// func NewMailboxType(settings Settings, config *configuration.Config) *MailboxType {
// 	return &MailboxType{
// 		settings: settings,
// 		config:   config,
// 	}
// }

// func (p *MailboxType) Create(owner ActorRef, system ActorSystem) (queue MessageQueue, err error) {
// 	return
// }

// type Mailbox struct {
// 	actor *ActorCell

// 	messageQueue MessageQueue
// }

// func NewMailbox(messageQueue MessageQueue) *Mailbox {
// 	return &Mailbox{
// 		messageQueue: messageQueue,
// 	}
// }

// func (p *Mailbox) Dispatcher() MessageDispatcher {
// 	return p.actor.Dispatcher()
// }

// func (p *Mailbox) SetActor(actorCell *ActorCell) {
// 	p.actor = actorCell
// }

// func (p *Mailbox) Run() {

// }

// func (p *Mailbox) isClosed() bool {
// 	return false
// }

// func (p *Mailbox) hasMessage() bool {
// 	return false
// }

// func (p *Mailbox) numberOfMessages() int {
// 	return 0
// }
