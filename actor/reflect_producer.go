package actor

import (
	"errors"
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
	errorType             = reflect.TypeOf((*error)(nil)).Elem()
)

type _ReflectProducer struct {
	typ  reflect.Type
	args []interface{}
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

	if !isCombinedUnTypedActor(p.typ) {
		err = ErrNoActorBaseCombind
		return
	}

	p.args = args
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

	untypedActor := NewUntypedActor(receiver.Receive)

	combinedUnTypedActor(val, untypedActor)

	actor = untypedActor.ActorBase

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

func isCombinedUnTypedActor(v reflect.Type) (isCombined bool) {
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Type == unTypedActorPtrType {
			return true
		}
	}

	return
}

func combinedUnTypedActor(val reflect.Value, unTypedActor *UntypedActor) (err error) {

	typ := val.Elem().Type()

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Type == unTypedActorPtrType {
			actorVal := reflect.ValueOf(unTypedActor)
			val.Elem().Field(i).Set(actorVal)
			return
		}
	}

	err = ErrNoActorBaseCombind
	return
}

func createInstanceByType(typ reflect.Type, args ...interface{}) (v reflect.Value, err error) {
	typVal := reflect.New(typ)

	if !typVal.IsValid() {
		err = ErrCreateInstanceFailure
		return
	}

	if err = initInstance(typVal, args...); err != nil {
		return
	}

	v = typVal

	return
}
