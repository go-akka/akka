package akka

type LoggingBus interface {
	EventBus

	SetLogLevel(level LogLevel)
	LogLevel() LogLevel

	StartStdoutLogger(config *Settings)
}
