package akka

type ActorCell interface {
	Self() ActorRef
	Mailbox() Mailbox

	SystemInvoke(message SystemMessage) (wasHandled bool, err error)
	Invoke(envelop Envelope) (wasHandled bool, err error)
	Dispatcher() MessageDispatcher
}
