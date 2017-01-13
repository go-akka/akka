package akka

type LogLevel int

type EventStream interface {
	StartUnsubscriber()
	Subscribe(subscriber ActorRef, channel interface{}) bool
	Unsubscribe(subscriber ActorRef, channels ...interface{}) bool
}
