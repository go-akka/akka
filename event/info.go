package event

import (
	"fmt"
	"github.com/go-akka/akka"
)

type Info struct {
	*LogEventBase
}

func NewInfoEvent(logSource string, logClass interface{}, message interface{}) akka.LogEvent {
	eventBase := newLogEventBase(logSource, logClass, message)
	event := &Info{
		LogEventBase: eventBase,
	}

	return event
}

func (p *Info) LogLevel() akka.LogLevel {
	return akka.InfoLevel
}

func (p *Info) String() string {
	return fmt.Sprintf("[%s][%s][%s] [%v]", p.LogLevel(), p.Timestamp(), p.LogSource(), p.Message())
}
