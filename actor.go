package akka

type Receiver interface {
	Receive(message interface{}) (handled bool, err error)
}

type ContextReceiver interface {
	Receive(context ActorContext, message interface{}) (handled bool, err error)
}

type Actor interface {
	Receiver
}

type MinimalActor interface {
	ContextReceiver
}

type PreStarter interface {
	PreStart() (err error)
}

type ContextPreStarter interface {
	PreStart(context ActorContext) (err error)
}

type Constructer interface {
	Construct() error
}

type InitFunc func() error
