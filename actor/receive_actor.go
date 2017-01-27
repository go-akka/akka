package actor

import (
	"github.com/go-akka/akka"
	"reflect"
)

type ReceiveActor struct {
	*ActorBase
	receiveFuns map[string]interface{}
	initFn      akka.InitFunc
}

func NewReceiveActor(actor akka.Actor, initFn akka.InitFunc) *ReceiveActor {

	receiveActor := &ReceiveActor{
		receiveFuns: make(map[string]interface{}),
	}

	// receiveActor.ActorBase = NewActorBase(receiveActor.Receive, actor)

	return receiveActor
}

func (p *ReceiveActor) Init() error {
	if p.initFn != nil {
		return p.initFn()
	}
	return nil
}

func (p *ReceiveActor) SetActorBase(actorBase *ActorBase) {
	p.ActorBase = actorBase
}

func (p *ReceiveActor) Receive(message interface{}) (handled bool, err error) {
	msgType := reflect.TypeOf(message)
	if fn, exist := p.receiveFuns[msgType.String()]; exist {
		handled = true

		retVals := reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(message)})
		if len(retVals) == 0 {
			return
		}

		lastVal := retVals[len(retVals)-1]
		if lastVal.Type() == errorType {
			err = lastVal.Interface().(error)
			return
		}

		return
	}
	return
}

func (p *ReceiveActor) SmartReceive(fn interface{}) (err error) {
	if fn == nil {
		err = ErrSmartReceiveArgIsNil
		return
	}

	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		err = ErrSmartReceiveShouldBeFunc
		return
	}

	if fnType.NumIn() != 1 {
		err = ErrWrongArgNumForSmartReceiveFunc
		return
	}

	argType := fnType.In(0)

	// TODO:
	// 1. add concurrent support
	// 2. Become and Unbecome ?
	p.receiveFuns[argType.String()] = fn

	return
}
