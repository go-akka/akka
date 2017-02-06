package event

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/fatih/color"
	"github.com/go-akka/akka"
)

var (
	stdDebugColor   = color.New(90)
	stdInfoColor    = color.New(color.FgWhite)
	stdWarningColor = color.New(color.FgYellow)
	stdErrorColor   = color.New(color.FgRed)
)

var (
	StandardOutLoggerInstance = NewStandardOutLogger(!color.NoColor, 100)
	StandardOutLoggerType     = reflect.TypeOf((*StandardOutLogger)(nil)).Elem()
)

type StandardOutLogger struct {
	*akka.MinimalActorRef

	UseColor bool

	eventChan chan interface{}
}

func NewStandardOutLogger(useColor bool, bufSize int) *StandardOutLogger {
	path := akka.NewRootActorPath(akka.NewAddress("akka", "all-systems", "", 0), "/StandardOutLogger")
	logger := &StandardOutLogger{
		MinimalActorRef: akka.NewMinimalActorRef(path, nil),
		UseColor:        useColor,
		eventChan:       make(chan interface{}, bufSize),
	}

	logger.start()

	return logger
}

func (p *StandardOutLogger) start() {
	go func() {
		for {
			select {
			case message, ok := <-p.eventChan:
				{
					if !ok {
						return
					}

					event, ok := message.(akka.LogEvent)
					if ok {
						p.printLogEvent(event)
					} else {
						fmt.Println(message)
					}
				}
			}
		}

	}()
}

func (s *StandardOutLogger) Provider() akka.ActorRefProvider {
	panic("This logger does not provide.")
}

func (p *StandardOutLogger) Tell(message interface{}, sender ...akka.ActorRef) (err error) {
	if message == nil {
		errors.New("The message to log must not be null.")
	}

	p.eventChan <- message

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
