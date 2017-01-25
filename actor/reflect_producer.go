package actor

import (
	"errors"
	"fmt"
	"github.com/go-akka/akka"
	"reflect"
)

type ActorBaseInitFunc func(instance interface{}) (err error)
type instanceCreateFunc func(val reflect.Value, args ...interface{}) (v reflect.Value, err error)

var (
	ErrCreateInstanceFailure = errors.New("create instance failure")
)

var (
	actorBaseInitFuncType = reflect.TypeOf((*ActorBaseInitFunc)(nil)).Elem()
	actorBasePtrType      = reflect.TypeOf((*ActorBase)(nil))
	unTypedActorPtrType   = reflect.TypeOf((*UntypedActor)(nil))
	receiveActorPtrType   = reflect.TypeOf((*ReceiveActor)(nil))
	errorType             = reflect.TypeOf((*error)(nil)).Elem()
)

type _ReflectProducer struct {
	typ      reflect.Type
	args     []interface{}
	baseType reflect.Type
}

func _NewReflectProducer(v interface{}, args ...interface{}) (producer IndirectActorProducer, err error) {
	p := _ReflectProducer{}
	if err = p.Init(v, args...); err != nil {
		return
	}
	producer = &p
	return
}

func (p *_ReflectProducer) Init(v interface{}, args ...interface{}) (err error) {
	if typ, ok := v.(reflect.Type); ok {
		p.typ = typ
	} else {
		p.typ = reflect.TypeOf(v).Elem()
	}

	if isCombined(p.typ, unTypedActorPtrType) {
		p.args = args
		p.baseType = unTypedActorPtrType
		return
	} else if isCombined(p.typ, receiveActorPtrType) {
		p.args = args
		p.baseType = receiveActorPtrType
		return
	}

	err = ErrNoUntypedActorOrReceiveActorCombind

	return
}

func (p *_ReflectProducer) Produce() (actor akka.Actor, err error) {

	var val reflect.Value
	if val, err = createInstanceByType(p.typ, p.args...); err != nil {
		return
	}

	if !val.IsValid() ||
		val.IsNil() {
		err = ErrCreateInstanceFailure
		return
	}

	receiver := val.Interface().(akka.Receiver)

	if p.baseType == unTypedActorPtrType {
		untypedActor := NewUntypedActor(receiver)
		combine(val, unTypedActorPtrType, untypedActor)
		actor = untypedActor.ActorBase
	} else if p.baseType == receiveActorPtrType {
		receiveActor := NewReceiveActor(receiver)
		combine(val, receiveActorPtrType, receiveActor)
		actor = receiveActor.ActorBase
	}

	if err = initInstance(val, p.args...); err != nil {
		return
	}

	return
}

func (p *_ReflectProducer) ActorType() reflect.Type {
	return p.typ
}

func initInstance(val reflect.Value, args ...interface{}) (err error) {

	methodVal := val.MethodByName("Init")

	if !methodVal.IsValid() {
		return
	}

	outNum := methodVal.Type().NumOut()
	if outNum > 1 {
		err = ErrBadActorInitFuncOutNumber
		return
	}

	if outNum == 1 {
		if methodVal.Type().Out(0) != errorType {
			err = ErrBadActorInitFuncOutType
			return
		}
	}

	var valArgs []reflect.Value
	for _, arg := range args {
		valArgs = append(valArgs, reflect.ValueOf(arg))
	}

	fnRetVals := methodVal.Call(valArgs)

	if len(fnRetVals) > 0 &&
		fnRetVals[0].IsValid() &&
		!fnRetVals[0].IsNil() {
		err = fnRetVals[0].Interface().(error)
		return
	}

	return
}

func isCombinedActorBase(v reflect.Type) (isCombined bool) {
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Type == actorBasePtrType {
			return true
		}
	}

	return
}

func isCombined(v reflect.Type, combineType reflect.Type) bool {
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Type == combineType {
			return true
		}
	}

	return false
}

func combine(val reflect.Value, combineType reflect.Type, combineValue interface{}) (err error) {

	typ := val.Elem().Type()

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Type == combineType {
			actorVal := reflect.ValueOf(combineValue)
			val.Elem().Field(i).Set(actorVal)
			return
		}
	}

	err = fmt.Errorf("struct should combine %s", combineType.String())
	return
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
