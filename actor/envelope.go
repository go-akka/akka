package actor

type Envelope struct {
	Message interface{}
	Sender  ActorRef
}
