package akka

type LogEvent interface {
	Level() LogLevel
	Message() interface{}
}
