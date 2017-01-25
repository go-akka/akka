package akka

type Receiver interface {
	Receive(message interface{}) (handled bool, err error)
}

type Actor interface {
	Receiver
}

type AutoReceivedMessage interface {
	AutoReceivedMessage()
}

type PreStarter interface {
	PreStart() (err error)
}
