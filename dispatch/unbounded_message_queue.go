package dispatch

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/pkg/lfqueue"
)

type UnboundedMessageQueue struct {
	queue *lfqueue.LockfreeQueue
}

func NewUnboundedMessageQueue() akka.MessageQueue {
	return &UnboundedMessageQueue{
		queue: lfqueue.NewLockfreeQueue(),
	}
}

func (p *UnboundedMessageQueue) Enqueue(receiver akka.ActorRef, envelope akka.Envelope) (err error) {
	p.queue.Push(envelope)
	return
}

func (p *UnboundedMessageQueue) Dequeue() (envelope akka.Envelope, ok bool) {
	v := p.queue.Pop()

	if v == nil {
		return
	}

	return v.(akka.Envelope), true
}

func (p *UnboundedMessageQueue) CleanUp(owner akka.ActorRef, deadLetters akka.MessageQueue) (err error) {

	for {
		msg, ok := p.Dequeue()
		if !ok {
			return
		}

		deadLetters.Enqueue(owner, msg)
	}

	return
}

func (p *UnboundedMessageQueue) NumberOfMessages() int {
	return int(p.queue.Size())
}

func (p *UnboundedMessageQueue) HasMessages() bool {
	return !p.queue.IsEmpty()
}
