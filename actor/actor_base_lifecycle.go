package actor

func (p *ActorBase) AroundPreReStart(cause error, message interface{}) {
	p.PreRestart(cause, message)
}

func (p *ActorBase) AroundPreStart() (err error) {
	return p.PreStart()
}

func (p *ActorBase) PreStart() (err error) {
	return
}

func (p *ActorBase) AroundPostRestart(cause error, message interface{}) {
	p.PostRestart(cause)
}

func (p *ActorBase) PreRestart(cause error, message interface{}) {
	for _, child := range p.Context().Children() {
		p.Context().Unwatch(child)
		p.Context().StopActor(child)
	}
	p.PostStop()
}

func (p *ActorBase) PostRestart(cause error) {
	p.PreStart()
}

func (p *ActorBase) AroundPostStop() {
	p.PostStop()
}

func (p *ActorBase) PostStop() (err error) {
	return
}
