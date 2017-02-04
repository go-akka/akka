package event

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-akka/akka"
	"reflect"
)

var (
	stdDebugColor   = color.New(90)
	stdInfoColor    = color.New(color.FgWhite)
	stdWarningColor = color.New(color.FgYellow)
	stdErrorColor   = color.New(color.FgRed)
)

var (
	StandardOutLoggerInstance = NewStandardOutLogger(true)
	StandardOutLoggerType     = reflect.TypeOf((*StandardOutLogger)(nil)).Elem()
)

type StandardOutLogger struct {
	*akka.MinimalActorRef

	UseColor bool
}

func NewStandardOutLogger(useColor bool) *StandardOutLogger {
	path := akka.NewRootActorPath(akka.NewAddress("akka", "all-systems", "", 0), "/StandardOutLogger")
	return &StandardOutLogger{
		MinimalActorRef: akka.NewMinimalActorRef(path, nil),
		UseColor:        useColor,
	}
}

func (s *StandardOutLogger) Provider() akka.ActorRefProvider {
	panic("This logger does not provide.")
}

func (p *StandardOutLogger) Tell(message interface{}, sender ...akka.ActorRef) (err error) {
	if message == nil {
		errors.New("The message to log must not be null.")
	}

	event, ok := message.(akka.LogEvent)
	if ok {
		p.printLogEvent(event)
	} else {
		fmt.Println(message)
	}

	return
}

func (p *StandardOutLogger) loglevelColor(level akka.LogLevel) *color.Color {
	switch level {
	case akka.DebugLevel:
		{
			return stdDebugColor
		}
	case akka.InfoLevel:
		{
			return stdInfoColor
		}
	case akka.WarningLevel:
		{
			return stdWarningColor
		}
	case akka.ErrorLevel:
		{
			return stdErrorColor
		}
	}

	return stdInfoColor
}

func (p *StandardOutLogger) printLogEvent(event akka.LogEvent) {
	if p.UseColor {
		colour := p.loglevelColor(event.LogLevel())
		colour.Println(event)
		return
	}

	fmt.Println(event)
}
