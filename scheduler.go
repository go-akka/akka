package akka

import (
	"time"
)

type CanTell interface {
	Tell(message interface{}, sender ...ActorRef) error
}

type TimeProvider interface {
	Now() time.Time
	MonotonicClock() time.Duration
	HighResMonotonicClock() time.Duration
}

type Cancelable interface {
	IsCancellationRequested() bool
	CancelAfter(delay time.Duration)
	Cancel(throwOnFirstException bool) (err error)
}

type TellScheduler interface {
	ScheduleTellOnce(delay time.Duration, receiver CanTell, message interface{}, sender ActorRef, cancelable Cancelable)
	ScheduleRepeatedly(delay time.Duration, interval time.Duration, receiver CanTell, message interface{}, sender ActorRef, cancelable Cancelable)
}

type ActionScheduler interface {
	ScheduleOnce(delay time.Duration, action Action, cancelable Cancelable)
	ScheduleRepeatedly(initialDelay time.Duration, interval time.Duration, action Action, cancelable Cancelable)
}

type AdvancedScheduler interface {
	ActionScheduler
}

type Scheduler interface {
	Advanced() AdvancedScheduler
}
