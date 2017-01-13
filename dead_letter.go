package akka

type DeadLetter struct {
	message   interface{}
	sender    ActorRef
	recipient ActorRef
}

func NewDeadLetter(message interface{}, sender ActorRef, recipient ActorRef) DeadLetter {
	return DeadLetter{
		message:   message,
		sender:    sender,
		recipient: recipient,
	}
}
