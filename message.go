package akka

type SystemMessage interface {
	SystemMessage()
}

type AutoReceivedMessage interface {
	AutoReceivedMessage()
}

type UnhandledMessage struct {
	Message   interface{}
	Sender    ActorRef
	Recipient ActorRef
}
