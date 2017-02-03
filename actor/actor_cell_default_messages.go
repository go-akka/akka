package actor

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/dispatch/sysmsg"
)

func (p *ActorCell) Invoke(msg akka.Envelope) (wasHandled bool, err error) {

	p.sender = msg.Sender

	switch message := msg.Message.(type) {
	case akka.AutoReceivedMessage:
		{
			return p.AutoReceiveMessage(msg)
		}
	default:
		return p.ReceiveMessage(message)
	}
}

func (p *ActorCell) SystemInvoke(msg akka.SystemMessage) (wasHandled bool, err error) {
	switch v := msg.(type) {
	case *sysmsg.Create:
		{
			p.create(v.Failure)
		}
	case *sysmsg.Terminate:
		{
			p.terminate()
		}
	}
	return
}

func (p *ActorCell) ReceiveMessage(message interface{}) (wasHandled bool, err error) {
	fn, exist := p.behaviorStack.Current()
	if !exist {
		// TODO:Create
		// retrun error
	}
	return p.actor.AroundReceive(fn, message)
}

func (p *ActorCell) AutoReceiveMessage(msg akka.Envelope) (wasHandled bool, err error) {
	switch val := msg.Message.(type) {
	case *akka.Terminated:
		{
			p.ReceivedTerminated(val)
		}
	case *akka.AddressTerminated:
		{

		}
	case *akka.Kill:
		{

		}
	case *akka.PoisonPill:
		{
			p.self.Stop()
		}
	case *akka.ActorSelectionMessage:
		{

		}
	case *akka.Identify:
		{
			//p.Sender().Tell(message, sender)
		}
	}
	return
}

func (p *ActorCell) create(failure error) {

	if failure != nil {
		panic(failure)
	}

	created, err := p.props.NewActor()
	if err != nil {
		panic(err)
	}

	actor := NewActorBase(created, p)

	if setter, ok := created.(actorBaseSetter); ok {
		setter.SetActorBase(actor)
	}

	if constructer, ok := created.(akka.Constructer); ok {
		constructer.Construct()
	}

	if err != nil {
		panic(err)
	}

	p.actor = actor

	if err = actor.AroundPreStart(); err != nil {
		panic(err)
	}

	p.actor.Become(p.actor.Receive, false)
}

func (p *ActorCell) matchSender(envelope akka.Envelope) akka.ActorRef {
	sender := envelope.Sender
	if sender == nil {
		sender = p.system.deadletters
	}
	return sender
}
