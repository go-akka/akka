package actor

import (
	"errors"
)

var (
	ErrNoActorProducerSpecified  = errors.New("No actor producer specified")
	ErrNoActorBaseCombind        = errors.New("actor should combine *akka.ActorBase")
	ErrBadActorInitFuncOutNumber = errors.New("the actor init func return's number should be 0 or 1,the type should be void or error")
	ErrBadActorInitFuncOutType   = errors.New("the actor init func return's should be void or error")
)
