package event

import (
	"reflect"
	"time"
)

type LogEventBase struct {
	timestamp  time.Time
	logSource  string
	logClass   reflect.Type
	message    interface{}
	stacktrace string
}

func newLogEventBase(logSource string, logClass, message interface{}) *LogEventBase {
	return &LogEventBase{
		timestamp: time.Now(),
		logSource: logSource,
		logClass:  reflect.TypeOf(logClass),
		message:   message,
	}

	//TODO add stacktrace
}

func (p *LogEventBase) Timestamp() time.Time {
	return p.timestamp
}

func (p *LogEventBase) LogSource() string {
	return p.logSource
}

func (p *LogEventBase) LogClass() reflect.Type {
	return p.logClass
}

func (p *LogEventBase) Stacktrace() string {
	return p.stacktrace
}

func (p *LogEventBase) Message() interface{} {
	return p.message
}
