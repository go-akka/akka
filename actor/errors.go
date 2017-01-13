package actor

import (
	"errors"
)

var (
	ErrNoActorProducerSpecified            = errors.New("No actor producer specified")
	ErrNoUntypedActorCombind               = errors.New("actor should combine *actor.UntypedActor")
	ErrNoReceiveActorCombind               = errors.New("actor should combine *actor.ReceiveActor")
	ErrNoUntypedActorOrReceiveActorCombind = errors.New("actor should combine *actor.UntypedActor or *actor.ReceiveActor")
	ErrBadActorInitFuncOutNumber           = errors.New("the actor init func return's number should be 0 or 1,the type should be void or error")
	ErrBadActorInitFuncOutType             = errors.New("the actor init func return's should be void or error")
	ErrSmartReceiveShouldBeFunc            = errors.New("smart receiver for args should be a func")
	ErrWrongArgNumForSmartReceiveFunc      = errors.New("wrong arg num for smart receive func")
	ErrSmartReceiveArgIsNil                = errors.New("smart receive arg should not be nil")
)
