package actor

func (p *ActorBase) PreStart() (err error) {
	return
}

func (p *ActorBase) PostStop() (err error) {
	return
}

func (p *ActorBase) PreRestart(err error, message interface{}) {
	return
}

func (p *ActorBase) PostRestart(err error) {
	return
}
