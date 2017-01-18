package akka

type Classifier interface {
	Classify(event interface{}, classifier interface{}) bool
	GetClassifier(event interface{}) (classifier interface{})
}

type Publisher interface {
	PublishToSubscriber(event interface{}, subscriber interface{})
}

type EventBus interface {
	TSubscribe(subscriber interface{}, classifier interface{}) bool
	TUnsubscribe(subscriber interface{}, classifiers ...interface{}) bool
	Publish(event interface{})
}
