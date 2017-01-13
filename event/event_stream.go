package event

import (
	"github.com/go-akka/akka"
	"reflect"
)

type EventStream struct {
	akka.LoggingBus

	system akka.ActorSystem
}

func NewEventStream(sys akka.ActorSystem, debug bool) akka.EventStream {
	eventStream := &EventStream{}

	eventBus := NewSubchannelClassification(eventStream, eventStream)
	eventStream.LoggingBus = NewLoggingBus(eventBus)

	return eventStream
}

func (p *EventStream) StartUnsubscriber() {
}

func (p *EventStream) Publish(event interface{}, subscriber interface{}) {

	sub, ok := subscriber.(akka.ActorRef)
	if !ok {
		return
	}

	if p.system == nil {
		p.Unsubscribe(sub)
		return
	}

	sub.Tell(event, akka.NoSender)
}

func (p *EventStream) Classify(event interface{}) interface{} {
	return reflect.TypeOf(event)
}

func (p *EventStream) Subscribe(subscriber akka.ActorRef, channel interface{}) bool {
	return p.LoggingBus.Subscribe(subscriber, channel)
}

func (p *EventStream) Unsubscribe(subscriber akka.ActorRef, channels ...interface{}) bool {
	return p.LoggingBus.Unsubscribe(subscriber, channels...)
}
