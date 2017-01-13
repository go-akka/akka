package akka

type Envelope struct {
	Message interface{}
	Sender  ActorRef
}
