package akka

type Classifier interface {
	Classify(event interface{}, classifier interface{}) bool
	GetClassifier(event interface{})
}

type Publisher interface {
	Publish(event interface{}, subscriber interface{})
}

type EventBus interface {
	Subscribe(subscriber interface{}, classifier interface{}) bool
	Unsubscribe(subscriber interface{}, classifiers ...interface{}) bool
	Publish(event interface{})
}
