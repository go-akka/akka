package akka

type EventStream interface {
	LoggingBus

	StartUnsubscriber()
	Subscribe(subscriber ActorRef, channel interface{}) bool
	Unsubscribe(subscriber ActorRef, channels ...interface{}) bool
	PublishToSubscriber(event interface{}, subscriber interface{})
}
