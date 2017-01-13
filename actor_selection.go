package akka

type ActorSelection struct {
	context ActorContext
}

func (p *ActorSelection) Tell(message interface{}, sender ActorRef) (err error) {
	return
}

func (p *ActorSelection) Forward(message interface{}) {
	p.Tell(message, p.context.Sender())
}
