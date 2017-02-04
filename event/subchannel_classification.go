package event

import (
	"github.com/go-akka/akka"
	"sync"
)

type SubchannelClassification struct {
	publisher  akka.Publisher
	classifier akka.Classifier
	classes    map[interface{}][]interface{}

	locker sync.Mutex
}

func NewSubchannelClassification(publisher akka.Publisher, classifier akka.Classifier) *SubchannelClassification {
	return &SubchannelClassification{
		publisher:  publisher,
		classifier: classifier,
		classes:    make(map[interface{}][]interface{}),
	}
}

func (p *SubchannelClassification) Publish(event interface{}) {

	// TODO: this is a temp solution

	c := p.classifier.GetClassifier(event)

	subscribers, exist := p.classes[c]
	if !exist {
		return
	}

	for i := 0; i < len(subscribers); i++ {
		p.publisher.PublishToSubscriber(event, subscribers[i])
	}
}

func (p *SubchannelClassification) TSubscribe(subscriber interface{}, class interface{}) bool {

	// TODO: this is a temp solution

	p.locker.Lock()
	defer p.locker.Unlock()

	oldsubs := p.classes[class]
	for i := 0; i < len(oldsubs); i++ {
		if oldsubs[i] == subscriber {
			return false
		}
	}

	p.classes[class] = append(p.classes[class], subscriber)

	return true
}

func (p *SubchannelClassification) TUnsubscribe(subscriber interface{}, classes ...interface{}) bool {

	// TODO: this is a temp solution

	p.locker.Lock()
	defer p.locker.Unlock()

	if len(classes) == 0 {
		for k, _ := range p.classes {
			classes = append(classes, k)
		}
	}

	for i := 0; i < len(classes); i++ {
		oldsubs := p.classes[classes[i]]
		for j := 0; j < len(oldsubs); j++ {
			if oldsubs[j] == subscriber {
				newSubs := oldsubs[0:j]
				newSubs = append(newSubs, oldsubs[j+1:])
				p.classes[classes[i]] = newSubs
				break
			}
		}
	}

	return false
}
