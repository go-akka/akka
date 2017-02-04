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
	Debug(foramt string, args ...interface{})
	Error(foramt string, args ...interface{})
	Info(foramt string, args ...interface{})
	Warning(foramt string, args ...interface{})
	Log(Level LogLevel, foramt string, args ...interface{})

	IsDebugEnabled() bool
	IsErrorEnabled() bool
	IsInfoEnabled() bool
	IsWarningEnabled() bool
}

type LogMessageFormatter interface {
	Format(format string, args ...interface{})
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
