package actor

type MailboxType interface {
	Create(owner []ActorRef, system []ActorSystem) (queue MessageQueue, err error)
}
