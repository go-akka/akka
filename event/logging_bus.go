package event

import (
	"github.com/go-akka/akka"
)

type LoggingBus struct {
	akka.EventBus

	loggers  []akka.ActorRef
	logLevel akka.LogLevel
}

func NewLoggingBus(classification akka.EventBus) *LoggingBus {
	return &LoggingBus{
		EventBus: classification,
	}
}

func (p *LoggingBus) SetLogLevel(level akka.LogLevel) {
}

func (p *LoggingBus) LogLevel() akka.LogLevel {
	return p.logLevel
}

func (p *LoggingBus) StartStdoutLogger(config *akka.Settings) {
}
