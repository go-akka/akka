package akka

import (
	"strings"
	"time"
)

type LogEvent interface {
	Timestamp() time.Time
	Stacktrace() string
	Message() interface{}
	LogLevel() LogLevel
}

type LoggingAdapter interface {
	LoggingFilter

	IsEnabled(level LogLevel) bool

	Debug(format string, args ...interface{})
	Error(cause error, format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Log(Level LogLevel, format string, args ...interface{})
}

type LoggingFilter interface {
	IsDebugEnabled() bool
	IsErrorEnabled() bool
	IsInfoEnabled() bool
	IsWarningEnabled() bool
}

type LogMessageFormatter interface {
	Format(format string, args ...interface{}) string
}

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

func AllLogLevels() []LogLevel {
	return []LogLevel{DebugLevel, InfoLevel, WarningLevel, ErrorLevel}
}

func (p LogLevel) String() string {
	switch p {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	}
	return "UNKNOWN"
}

func LogLevelFor(level string) LogLevel {
	lv := strings.ToUpper(level)

	switch lv {
	case "DEBUG":
		return DebugLevel
	case "INFO":
		return InfoLevel
	case "WARN":
		return WarningLevel
	case "ERROR":
		return ErrorLevel
	default:
		panic("Unknown LogLevel: " + lv + ". Valid values are: DEBUG, INFO, WARN, ERROR")
	}
}
