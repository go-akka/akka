package akka

type UnhandledMessage struct {
	Message   interface{}
	Sender    ActorRef
	Recipient ActorRef
}

type SystemMessage interface {
	SystemMessage()
}
