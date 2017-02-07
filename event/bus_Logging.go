package event

import (
	"github.com/go-akka/akka"
	"reflect"
)

type BusLogging struct {
	bus                 akka.LoggingBus
	logClass            reflect.Type
	logSource           string
	logMessageFormatter akka.LogMessageFormatter

	isErrorEnabled, isWarningEnabled, isInfoEnabled, isDebugEnabled bool
}

func NewBusLogging(bus akka.LoggingBus, logSource string, logClass reflect.Type, logMessageFormatter akka.LogMessageFormatter) akka.LoggingAdapter {

	bugLogging := &BusLogging{
		bus:                 bus,
		logSource:           logSource,
		logClass:            logClass,
		logMessageFormatter: logMessageFormatter,
		isErrorEnabled:      bus.LogLevel() <= akka.ErrorLevel,
		isWarningEnabled:    bus.LogLevel() <= akka.WarningLevel,
		isInfoEnabled:       bus.LogLevel() <= akka.InfoLevel,
		isDebugEnabled:      bus.LogLevel() <= akka.DebugLevel,
	}

	adapter := NewLoggingAdapter(bugLogging, logMessageFormatter)

	return adapter
}

func (p *BusLogging) IsDebugEnabled() bool {
	return p.isDebugEnabled
}

func (p *BusLogging) IsErrorEnabled() bool {
	return p.isErrorEnabled
}

func (p *BusLogging) IsInfoEnabled() bool {
	return p.isInfoEnabled
}

func (p *BusLogging) IsWarningEnabled() bool {
	return p.isWarningEnabled
}

func (p *BusLogging) NotifyError(cause error, message interface{}) {
	p.bus.Publish(NewErrorEvent(cause, p.logSource, p.logClass, message))
}

func (p *BusLogging) NotifyWarning(message interface{}) {
	p.bus.Publish(NewWarningEvent(p.logSource, p.logClass, message))
}

func (p *BusLogging) NotifyInfo(message interface{}) {
	p.bus.Publish(NewInfoEvent(p.logSource, p.logClass, message))
}

func (p *BusLogging) NotifyDebug(message interface{}) {
	p.bus.Publish(NewDebugEvent(p.logSource, p.logClass, message))
}
