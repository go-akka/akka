package akka

type SystemMessage interface {
	SystemMessage()
}

type UnhandledMessage struct {
	Message   interface{}
	Sender    ActorRef
	Recipient ActorRef
}
