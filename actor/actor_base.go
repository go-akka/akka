package actor

var emptyBehavior = func(_ interface{}) bool {
	return false
}

type ActorBase struct {
	ctx ActorContext
}

func (p *ActorBase) String() string {
	return ""
}

func (p *ActorBase) Context() ActorContext {
	return p.ctx
}

func (p *ActorBase) PreStart() (err error) {
	return
}

func (p *ActorBase) PostStop() (err error) {
	return
}

func (p *ActorBase) PreRestart() (err error) {
	return
}

func (p *ActorBase) PostRestart() (err error) {
	return
}

func (p *ActorBase) Receive(message interface{}) bool {
	return false
}

func (p *ActorBase) Sender() ActorRef {
	return p.ctx.Sender()
}

func (p *ActorBase) Self() ActorRef {
	return p.ctx.Self()
}

func (p *ActorBase) Become(receive Receive) (err error) {
	return p.ctx.Become(receive)
}

func (p *ActorBase) BecomeStacked(receive Receive) (err error) {
	return p.ctx.BecomeStacked(receive)
}

func (p *ActorBase) UnbecomeStacked() (err error) {
	return p.Context().UnbecomeStacked()
}

func (p *ActorBase) Unhandled(message interface{}) (err error) {
	return
}
