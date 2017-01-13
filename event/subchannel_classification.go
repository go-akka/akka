package event

import (
	"github.com/go-akka/akka"
	"github.com/orcaman/concurrent-map"
)

type SubchannelClassification struct {
	publisher  akka.Publisher
	classifier akka.Classifier
	classes    cmap.ConcurrentMap
}

func NewSubchannelClassification(publisher akka.Publisher, classifier akka.Classifier) *SubchannelClassification {
	return &SubchannelClassification{
		publisher:  publisher,
		classifier: classifier,
	}
}

func (p *SubchannelClassification) Publish(event interface{}) {
}

func (p *SubchannelClassification) Subscribe(subscriber interface{}, class interface{}) bool {
	return false
}

func (p *SubchannelClassification) Unsubscribe(subscriber interface{}, classes ...interface{}) bool {
	return false
}
