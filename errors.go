package akka

import (
	"errors"
)

var (
	ErrInvalidActorSystemName               = errors.New("invalid ActorSystem name, must contain only word characters (i.e. [a-zA-Z0-9] plus non-leading '-' or '_')")
	ErrBadNumOutOfCreateInstanceFunc        = errors.New("bad out number of instance create func (number: 1~2)")
	ErrBadNumInOfCreateInstanceFunc         = errors.New("bad in number of instance create func (number: >1), the first arg should be: func(instance interface{})(error)")
	ErrBadTypeOfCreateInstanceFuncSecondOut = errors.New("the second create instance func of out should be type of error")
	ErrActorBaseInitFuncNotCalled           = errors.New("the first args is func, you should call it first before you use")
	ErrBadTypeOfCreateInstanceFuncFirstIn   = errors.New("the first arg type of create instance func should be: func(instance interface{})(error)")
	ErrCreateActorRefProviderFailure        = errors.New("create actore ref provider failure")
	ErrBadTypeOfScheduler                   = errors.New("basd scheduler type")
	ErrTypeNotExistInClassLoader            = errors.New("type not in class loader")
)
