package event

import (
	"github.com/go-akka/akka"
)

type SubchannelClassification struct {
	publisher akka.Publisher
}

func NewSubchannelClassification(publisher akka.Publisher, classifier akka.Classifier) *SubchannelClassification {
	return &SubchannelClassification{}
}

func (p *SubchannelClassification) Publish(event interface{}) {

}

func (p *SubchannelClassification) Subscribe(subscriber interface{}, classifier interface{}) bool {
	return false
}

func (p *SubchannelClassification) Unsubscribe(subscriber interface{}, classifiers ...interface{}) bool {
	return false
}
