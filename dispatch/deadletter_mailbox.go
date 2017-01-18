package dispatch

import (
	"github.com/go-akka/akka"
)

type DeadLetterMailbox struct {
	*Mailbox

	deadLeaters akka.ActorRef
}

func NewDeadLetterMailbox() *DeadLetterMailbox {
	return &DeadLetterMailbox{}
}

func (p *DeadLetterMailbox) Enqueue(receiver akka.ActorRef, envelope akka.Envelope) (err error) {
	switch v := envelope.Message.(type) {
	case *akka.DeadLetter:
		{

		}
	default:
		p.deadLeaters.Tell(akka.NewDeadLetter(v, envelope.Sender, receiver), envelope.Sender)
	}

	return
}

func (p *DeadLetterMailbox) Dequeue() (envelope akka.Envelope, ok bool) {
	return
}

func (p *DeadLetterMailbox) CleanUp(owner akka.ActorRef, deadLetters akka.MessageQueue) (err error) {
	return
}
