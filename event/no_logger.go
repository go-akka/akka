package event

import (
	"github.com/go-akka/akka"
)

var (
	NoLoggerInstance akka.LoggingAdapter = &NoLogger{}
)

type NoLogger struct {
}

func (p *NoLogger) Debug(foramt string, args ...interface{}) {
	return
}

func (p *NoLogger) Error(foramt string, args ...interface{}) {
	return
}

func (p *NoLogger) Info(foramt string, args ...interface{}) {
	return
}

func (p *NoLogger) Warning(foramt string, args ...interface{}) {
	return
}

func (p *NoLogger) Log(Level akka.LogLevel, foramt string, args ...interface{}) {
	return
}

func (p *NoLogger) IsDebugEnabled() bool {
	return false
}

func (p *NoLogger) IsErrorEnabled() bool {
	return false
}

func (p *NoLogger) IsInfoEnabled() bool {
	return false
}

func (p *NoLogger) IsWarningEnabled() bool {
	return false
}
