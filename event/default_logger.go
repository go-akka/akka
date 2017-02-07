package event

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/pkg/class_loader"
)

func init() {
	class_loader.Default.Register((*DefaultLogger)(nil), "akka.event.default-logger")
}

var (
	_ akka.MinimalActor = (*DefaultLogger)(nil)
)

type DefaultLogger struct {
}

func (p *DefaultLogger) Receive(context akka.ActorContext, message interface{}) (wasHandled bool, err error) {
	if event, ok := message.(akka.LogEvent); ok {
		p.Print(event)
	}
	wasHandled = true
	return
}

func (p *DefaultLogger) Print(logEvent akka.LogEvent) {
	StandardOutLoggerInstance.printLogEvent(logEvent)
}
