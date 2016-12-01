package actor

type Actor interface {
	Receive(message interface{}) (unhandled bool)
}
