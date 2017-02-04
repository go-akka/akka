package event

import (
	"fmt"
	"github.com/go-akka/akka"
)

type Warning struct {
	*LogEventBase
}

func NewWarningEvent(logSource string, logClass interface{}, message interface{}) akka.LogEvent {
	eventBase := newLogEventBase(logSource, logClass, message)
	event := &Warning{
		LogEventBase: eventBase,
	}

	return event
}

func (p *Warning) LogLevel() akka.LogLevel {
	return akka.WarningLevel
}

func (p *Warning) String() string {
	return fmt.Sprintf("[%s][%s][%s] [%v]", p.LogLevel(), p.Timestamp(), p.LogSource(), p.Message())
}
