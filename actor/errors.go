package actor

import (
	"errors"
)

var (
	ErrCreateInstanceFailure                = errors.New("create instance failure")
	ErrNoActorBaseCombind                   = errors.New("actor should combine actor.ActorBase")
	ErrBadNumOutOfCreateInstanceFunc        = errors.New("bad out number of instance create func (number: 1~2)")
	ErrBadNumInOfCreateInstanceFunc         = errors.New("bad in number of instance create func (number: >1), the first arg should be: func(instance interface{})(error)")
	ErrBadTypeOfCreateInstanceFuncSecondOut = errors.New("the second create instance func of out should be type of error")
	ErrActorBaseInitFuncNotCalled           = errors.New("the first args is func, you should call it first before you use")
	ErrBadTypeOfCreateInstanceFuncFirstIn   = errors.New("the first arg type of create instance func should be: func(instance interface{})(error)")
	ErrBadActorInitFuncOutNumber            = errors.New("the actor init func return's number should be 0 or 1,the type should be void or error")
	ErrBadActorInitFuncOutType              = errors.New("the actor init func return's should be void or error")
)
