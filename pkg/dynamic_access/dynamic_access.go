package dynamic_access

import (
	"errors"
	"fmt"
	"github.com/go-akka/akka/pkg/class_loader"

	"reflect"
)

var (
	errorType = reflect.TypeOf((*error)(nil)).Elem()
)

var (
	ErrBadActorInitFuncOutNumber = errors.New("the actor init func return's number should be 0 or 1,the type should be void or error")
	ErrBadActorInitFuncOutType   = errors.New("the actor init func return's should be void or error")
	ErrCreateInstanceFailure     = errors.New("create instance failure")
	ErrTypeNotExistInClassLoader = errors.New("type not in class loader")
)

type DynamicAccess interface {
	CreateInstanceByType(typ reflect.Type, args ...interface{}) (ins interface{}, err error)
	CreateInstanceByName(name string, args ...interface{}) (ins interface{}, err error)
}

type ReflectiveDynamicAccess struct {
	classLoader class_loader.ClassLoader
}

func NewReflectiveDynamicAccess(classLoader class_loader.ClassLoader) DynamicAccess {
	return &ReflectiveDynamicAccess{
		classLoader: classLoader,
	}
}

func (p *ReflectiveDynamicAccess) CreateInstanceByType(typ reflect.Type, args ...interface{}) (ins interface{}, err error) {
	typVal := reflect.New(typ)

	if !typVal.IsValid() {
		err = ErrCreateInstanceFailure
		return
	}

	if err = p.constructInstance(typVal, args...); err != nil {
		return
	}

	ins = typVal.Interface()

	return
}

func (p *ReflectiveDynamicAccess) CreateInstanceByName(name string, args ...interface{}) (ins interface{}, err error) {
	if typ, exist := p.classLoader.ClassNameOf(name); !exist {
		err = fmt.Errorf("[ErrTypeNotExistInClassLoader] TypeName: %s", name)
		return
	} else {
		return p.CreateInstanceByType(typ, args...)
	}

	return
}

func (p *ReflectiveDynamicAccess) constructInstance(val reflect.Value, args ...interface{}) (err error) {

	methodVal := val.MethodByName("Construct")

	if !methodVal.IsValid() {
		return
	}

	numOut := methodVal.Type().NumOut()
	if numOut > 1 {
		err = ErrBadActorInitFuncOutNumber
		return
	}

	if numOut == 1 && methodVal.Type().Out(0) != errorType {
		err = ErrBadActorInitFuncOutType
		return
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
