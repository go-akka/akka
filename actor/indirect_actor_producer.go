package actor

import (
	"github.com/go-akka/akka"
	"reflect"
)

type IndirectActorProducer interface {
	Produce() (actor akka.Actor, err error)
	ActorType() reflect.Type
}

func _CreateProducer(v interface{}, args ...interface{}) (producer IndirectActorProducer, err error) {
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
	} else if producer, err = _NewReflectProducer(v, args...); err != nil {
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
