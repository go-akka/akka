package event

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-akka/akka"
)

var (
	Logging = &logging{}
)

type DefaultLogMessageFormatter struct {
}

func (p *DefaultLogMessageFormatter) Format(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

type logging struct {
}

func LogClassFor(level akka.LogLevel) reflect.Type {

	switch level {
	case akka.DebugLevel:
		return reflect.TypeOf((*Debug)(nil)).Elem()
	case akka.InfoLevel:
		return reflect.TypeOf((*Info)(nil)).Elem()
	case akka.WarningLevel:
		return reflect.TypeOf((*Warning)(nil)).Elem()
	case akka.ErrorLevel:
		return reflect.TypeOf((*Error)(nil)).Elem()
	}

	panic("Unknown LogLevel: " + strconv.Itoa(int(level)))
}

func (p *logging) GetLogger(context akka.ActorContext, logMessageFormatter ...akka.LogMessageFormatter) akka.LoggingAdapter {
	logSource := context.Self().String()
	logClass := context.Props().Type()

	var formatter akka.LogMessageFormatter
	if len(logMessageFormatter) == 0 {
		formatter = &DefaultLogMessageFormatter{}
	} else {
		formatter = logMessageFormatter[0]
	}

	return NewBusLogging(context.System().EventStream(), logSource, logClass, formatter)
}

func (p *logging) GetLoggerWithActorSystem(system akka.ActorSystem, logSource interface{}, logMessageFormatter ...akka.LogMessageFormatter) akka.LoggingAdapter {
	logSourceStr := ""

	if str, ok := logSource.(string); ok {
		logSourceStr = str
	} else {
		logSourceStr = reflect.TypeOf(logSource).String()
	}

	var formatter akka.LogMessageFormatter
	if len(logMessageFormatter) == 0 {
		formatter = &DefaultLogMessageFormatter{}
	} else {
		formatter = logMessageFormatter[0]
	}

	return NewBusLogging(system.EventStream(), fmt.Sprintf("%s(%s)", logSourceStr, system), reflect.TypeOf(system), formatter)
}
