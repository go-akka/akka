package event

import (
	"github.com/go-akka/akka"
	"strconv"
)

type LoggingAdapter struct {
	LogNotifier
	logMessageFormatter akka.LogMessageFormatter
}

func NewLoggingAdapter(notifier LogNotifier, logMessageFormatter akka.LogMessageFormatter) *LoggingAdapter {

	if logMessageFormatter == nil {
		panic(simpleName(logMessageFormatter) + "The message formatter must not be nil.")
	}

	if notifier == nil {
		panic(simpleName(notifier) + "The log notifier must not be nil.")
	}

	base := &LoggingAdapter{
		LogNotifier:         notifier,
		logMessageFormatter: logMessageFormatter,
	}

	return base
}

func (p *LoggingAdapter) Debug(format string, args ...interface{}) {

	if !p.IsDebugEnabled() {
		return
	}

	if len(args) == 0 {
		p.NotifyDebug(format)
	} else {
		p.NotifyDebug(LogMessage{p.logMessageFormatter, format, args})
	}
}

func (p *LoggingAdapter) Error(cause error, format string, args ...interface{}) {
	if !p.IsErrorEnabled() {
		return
	}

	if len(args) == 0 {
		p.NotifyError(cause, format)
	} else {
		p.NotifyError(cause, LogMessage{p.logMessageFormatter, format, args})
	}
}

func (p *LoggingAdapter) Info(format string, args ...interface{}) {
	if !p.IsInfoEnabled() {
		return
	}

	if len(args) == 0 {
		p.NotifyInfo(format)
	} else {
		p.NotifyInfo(LogMessage{p.logMessageFormatter, format, args})
	}
}

func (p *LoggingAdapter) Warning(format string, args ...interface{}) {
	if !p.IsWarningEnabled() {
		return
	}

	if len(args) == 0 {
		p.NotifyWarning(format)
	} else {
		p.NotifyWarning(LogMessage{p.logMessageFormatter, format, args})
	}
}

func (p *LoggingAdapter) Log(level akka.LogLevel, format string, args ...interface{}) {
	if len(args) == 0 {
		p.NotifyLog(level, format)
	} else {
		p.NotifyLog(level, LogMessage{p.logMessageFormatter, format, args})
	}
	return
}

func (p *LoggingAdapter) NotifyLog(logLevel akka.LogLevel, message interface{}) {
	switch logLevel {
	case akka.DebugLevel:
		if p.IsDebugEnabled() {
			p.NotifyDebug(message)
		}
	case akka.InfoLevel:
		if p.IsInfoEnabled() {
			p.NotifyInfo(message)
		}
	case akka.WarningLevel:
		if p.IsWarningEnabled() {
			p.NotifyWarning(message)
		}
	case akka.ErrorLevel:
		if p.IsErrorEnabled() {
			p.NotifyError(nil, message)
		}
	default:
		panic("Unknown LogLevel: " + strconv.Itoa(int(logLevel)) + ". Valid values are: DEBUG, INFO, WARN, ERROR")
	}
}

func (p *LoggingAdapter) IsEnabled(level akka.LogLevel) bool {

	switch level {
	case akka.DebugLevel:
		return p.IsDebugEnabled()
	case akka.InfoLevel:
		return p.IsInfoEnabled()
	case akka.WarningLevel:
		return p.IsWarningEnabled()
	case akka.ErrorLevel:
		return p.IsErrorEnabled()
	default:
		panic("Unknown LogLevel: " + strconv.Itoa(int(level)) + ". Valid values are: DEBUG, INFO, WARN, ERROR")
	}
}
