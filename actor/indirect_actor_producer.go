package actor

import (
	"reflect"
)

var (
	_ IndirectActorProducer = (*_CreatorFunctionConsumer)(nil)
	_ IndirectActorProducer = (*_ArgsReflectConstructor)(nil)
)

type IndirectActorProducer interface {
	Produce() (actor Actor, err error)
	ActorType() reflect.Type
}

func NewIndirectActorProducer(typ reflect.Type, args ...interface{}) (producer IndirectActorProducer, err error) {
	return
}

type _CreatorFunctionConsumer struct {
	typ     reflect.Type
	creator func() (Actor, error)
}

func (p *_CreatorFunctionConsumer) Init(typ reflect.Type, creator func() (Actor, error)) {
	p.typ = typ
	p.creator = creator
}

func (p *_CreatorFunctionConsumer) Produce() (actor Actor, err error) {
	return p.creator()
}

func (p *_CreatorFunctionConsumer) ActorType() reflect.Type {
	return p.typ
}

type _ArgsReflectConstructor struct {
	typ  reflect.Type
	args []interface{}
}

func (p *_ArgsReflectConstructor) Init(typ reflect.Type, args ...interface{}) {
	p.typ = typ
	p.args = args
	return
}

func (p *_ArgsReflectConstructor) Produce() (actor Actor, err error) {
	return
}

func (p *_ArgsReflectConstructor) ActorType() reflect.Type {
	return p.typ
}
