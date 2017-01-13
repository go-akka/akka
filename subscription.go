package akka

type Subscription struct {
	subscriber      interface{}
	unsubscriptions []interface{}
}

func NewSubscription(subscriber interface{}) *Subscription {
	return &Subscription{
		subscriber: subscriber,
	}
}

func (p *Subscription) Subscription(subscriber interface{}, unsubscriptions []interface{}) {
	p.subscriber = subscriber
	p.unsubscriptions = unsubscriptions
}

func (p *Subscription) Subscriber() interface{} {
	return p.subscriber
}

func (p *Subscription) Unsubscriptions() []interface{} {
	return p.unsubscriptions
}
