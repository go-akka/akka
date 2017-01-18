package akka

type MessageQueue interface {
	Enqueue(receiver ActorRef, envelope Envelope) (err error)
	Dequeue() (envelope Envelope, ok bool)
	CleanUp(owner ActorRef, deadLetters MessageQueue) (err error)

	NumberOfMessages() int
	HasMessages() bool
}

type QueueBasedMessageQueue interface {
	MessageQueue
	Queue() []Envelope
}
