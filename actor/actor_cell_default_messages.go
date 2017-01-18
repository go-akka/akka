package actor

import (
	"github.com/go-akka/akka"
)

func (p *ActorCell) Invoke(msg akka.Envelope) (wasHandled bool, err error) {
	switch message := msg.Message.(type) {
	case akka.AutoReceivedMessage:
		{
			return p.AutoReceiveMessage(msg)
		}
	default:
		return p.ReceiveMessage(message)
	}
}

func (p *ActorCell) SystemInvoke(msg akka.Envelope) (wasHandled bool, err error) {
	return
}

func (p *ActorCell) ReceiveMessage(message interface{}) (wasHandled bool, err error) {
	fn, exist := p.behaviorStack.Current()
	if !exist {
		// TODO:
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
