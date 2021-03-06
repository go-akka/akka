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
	eventStream := &EventStream{system: sys}

	eventBus := NewSubchannelClassification(eventStream, eventStream)
	eventStream.LoggingBus = NewLoggingBus(eventBus)

	return eventStream
}

func (p *EventStream) StartUnsubscriber() {
}

func (p *EventStream) PublishToSubscriber(event interface{}, subscriber interface{}) {
	sub, ok := subscriber.(akka.ActorRef)
	if !ok {
		return
	}

	if p.system == nil {
		p.Unsubscribe(sub)
		return
	}

	sub.Tell(event)
}

func (p *EventStream) GetClassifier(event interface{}) interface{} {
	t := reflect.TypeOf(event)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func (p *EventStream) Classify(event interface{}, classifier interface{}) bool {
	return p.GetClassifier(event) == classifier.(reflect.Type)
}

func (p *EventStream) Subscribe(subscriber akka.ActorRef, channel interface{}) bool {
	return p.LoggingBus.TSubscribe(subscriber, channel)
}

func (p *EventStream) Unsubscribe(subscriber akka.ActorRef, channels ...interface{}) bool {
	return p.LoggingBus.TUnsubscribe(subscriber, channels...)
}
