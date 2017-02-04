package event

import (
	"fmt"
	"github.com/go-akka/akka"
)

type Debug struct {
	*LogEventBase
}

func NewDebugEvent(logSource string, logClass interface{}, message interface{}) akka.LogEvent {
	eventBase := newLogEventBase(logSource, logClass, message)
	event := &Debug{
		LogEventBase: eventBase,
	}

	return event
}

func (p *Debug) LogLevel() akka.LogLevel {
	return akka.DebugLevel
}

func (p *Debug) String() string {
	return fmt.Sprintf("[%s][%s][%s] [%v]", p.LogLevel(), p.Timestamp(), p.LogSource(), p.Message())
}
