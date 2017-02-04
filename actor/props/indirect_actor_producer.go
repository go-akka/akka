package props

import (
	"errors"
	"github.com/go-akka/akka"
	"reflect"
)

var (
	ErrNoActorProducerSpecified = errors.New("No actor producer specified")
	ErrCreateInstanceFailure    = errors.New("create instance failure")
)

type IndirectActorProducer interface {
	Produce() (actor akka.Actor, err error)
	ActorType() reflect.Type
}

func createProducer(producerCreator ProducerCreatorFunc, v interface{}, args ...interface{}) (producer IndirectActorProducer, err error) {
	typ := reflect.TypeOf(v)

	if typ == nil {
		producer = &_DefaultProducer{}
		return
	}

	if _, ok := v.(IndirectActorProducer); ok {
		var obj interface{}
		if obj, err = createInstanceByType(typ, args...); err != nil {
			return
		}

		producer = obj.(IndirectActorProducer)

		return
	} else if producer, err = producerCreator(v, args...); err != nil {
		return
	}

	return
}

type _DefaultProducer struct {
}

func (p *_DefaultProducer) Produce() (actor akka.Actor, err error) {
	err = ErrNoActorProducerSpecified
	return
}

func (p *_DefaultProducer) ActorType() reflect.Type {
	return nil
}

func createInstanceByType(typ reflect.Type, args ...interface{}) (v reflect.Value, err error) {
	typVal := reflect.New(typ)

	if !typVal.IsValid() {
		err = ErrCreateInstanceFailure
		return
	}

	v = typVal

	return
}
