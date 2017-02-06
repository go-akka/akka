package event

type LogNotifier interface {
	IsDebugEnabled() bool
	IsErrorEnabled() bool
	IsInfoEnabled() bool
	IsWarningEnabled() bool

	NotifyError(cause error, message interface{})
	NotifyWarning(message interface{})
	NotifyInfo(message interface{})
	NotifyDebug(message interface{})
}
