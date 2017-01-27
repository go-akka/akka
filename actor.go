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

type Constructer interface {
	Construct() error
}

type InitFunc func() error
