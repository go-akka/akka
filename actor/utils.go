package actor

import (
	"reflect"
)

type instanceCreateFunc func(val reflect.Value, args ...interface{}) (v reflect.Value, err error)
type ActorBaseInitFunc func(instance interface{}) (err error)

const (
	_ActorInitFuncName = "Init"
)

var (
	errorType             = reflect.TypeOf((*error)(nil)).Elem()
	actorBaseInitFuncType = reflect.TypeOf((*ActorBaseInitFunc)(nil)).Elem()
)

func createInstance(v interface{}, args ...interface{}) (obj interface{}, err error) {

	vType := reflect.TypeOf(v)
	vVal := reflect.ValueOf(v)

	var createFunc instanceCreateFunc
	var retVal reflect.Value

	switch vType.Kind() {
	case reflect.Struct:
		{
			createFunc = createInstanceByType
		}
	case reflect.Func:
		{
			createFunc = createInstanceByFunc
		}
	default:
		err = ErrCreateInstanceFailure
		return
	}

	if retVal, err = createFunc(vVal, args...); err != nil {
		return
	}

	if retVal.IsValid() ||
		retVal.IsNil() {
		err = ErrCreateInstanceFailure
		return
	}

	return
}

func createInstanceByType(val reflect.Value, args ...interface{}) (v reflect.Value, err error) {

	typVal := reflect.New(val.Type())
	if !typVal.IsValid() {
		err = ErrCreateInstanceFailure
		return
	}

	if err = initInstanceActorBase(typVal); err != nil {
		return
	}

	if err = initInstance(typVal, args...); err != nil {
		return
	}

	v = typVal

	return
}

func createInstanceByFunc(val reflect.Value, args ...interface{}) (v reflect.Value, err error) {

	fnTyp := val.Type()

	if fnTyp.NumIn() == 0 {
		err = ErrBadNumInOfCreateInstanceFunc
		return
	}

	if !fnTyp.In(0).ConvertibleTo(actorBaseInitFuncType) {
		err = ErrBadTypeOfCreateInstanceFuncFirstIn
		return
	}

	if fnTyp.NumOut() == 0 || fnTyp.NumOut() > 2 {
		err = ErrBadNumOutOfCreateInstanceFunc
		return
	}

	if fnTyp.NumOut() == 2 && fnTyp.Out(1) != errorType {
		err = ErrBadTypeOfCreateInstanceFuncSecondOut
		return
	}

	var valArgs []reflect.Value

	initCalled := false

	var initFn ActorBaseInitFunc
	initFn = func(ins interface{}) (e error) {
		insVal := reflect.ValueOf(ins)
		e = initInstanceActorBase(insVal, args...)
		initCalled = true
		return
	}

	valArgs = append(valArgs, reflect.ValueOf(initFn))

	for _, arg := range args {
		valArgs = append(valArgs, reflect.ValueOf(arg))
	}

	fnRetVals := val.Call(valArgs)

	if initCalled == false {
		err = ErrActorBaseInitFuncNotCalled
		return
	}

	if len(fnRetVals) == 1 {
		v = fnRetVals[0]
		return
	} else {
		if fnRetVals[1].IsValid() &&
			fnRetVals[1].IsNil() {
			v = fnRetVals[0]
			return
		}

		err = fnRetVals[1].Interface().(error)
		return
	}

	return
}

func initInstance(val reflect.Value, args ...interface{}) (err error) {

	methodVal := val.MethodByName(_ActorInitFuncName)

	if methodVal.Type().NumOut() > 1 {
		err = ErrBadActorInitFuncOutNumber
		return
	}

	if methodVal.Type().Out(0) != errorType {
		err = ErrBadActorInitFuncOutType
		return
	}

	if !methodVal.IsValid() {
		return
	}

	var valArgs []reflect.Value
	for _, arg := range args {
		valArgs = append(valArgs, reflect.ValueOf(arg))
	}

	fnRetVals := val.Call(valArgs)

	if len(fnRetVals) > 0 &&
		fnRetVals[0].IsValid() &&
		!fnRetVals[0].IsNil() {
		err = fnRetVals[0].Interface().(error)
		return
	}

	return
}

func initInstanceActorBase(v reflect.Value, args ...interface{}) (err error) {
	if !isCombinedActorBase(v.Type()) {
		err = ErrNoActorBaseCombind
		return
	}

	return
}

func isCombinedActorBase(v reflect.Type) (isCombined bool) {
	return
}
