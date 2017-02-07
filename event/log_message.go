package event

import (
	"github.com/go-akka/akka"
)

type LogMessage struct {
	Formatter akka.LogMessageFormatter
	Format    string
	Args      []interface{}
}

func (p LogMessage) String() string {
	return p.Formatter.Format(p.Format, p.Args...)
}
