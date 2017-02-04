package event

import (
	"fmt"
	"github.com/go-akka/akka"
)

type Error struct {
	*LogEventBase
	cause error
}

func NewErrorEvent(cause error, logSource string, logClass interface{}, message interface{}) akka.LogEvent {
	eventBase := newLogEventBase(logSource, logClass, message)
	event := &Error{
		LogEventBase: eventBase,
		cause:        cause,
	}

	return event
}

func (p *Error) LogLevel() akka.LogLevel {
	return akka.ErrorLevel
}

func (p *Error) String() string {
	causeStr := "Unknown"
	if p.cause != nil {
		causeStr = p.cause.Error()
	}

	return fmt.Sprintf("[%s][%s][%s] [%v]\nCause: [%s]", p.LogLevel(), p.Timestamp(), p.LogSource(), p.Message(), causeStr)
}
