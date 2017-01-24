package akka

type UnhandledMessage struct {
	Message   interface{}
	Sender    ActorRef
	Recipient ActorRef
}

type SystemMessage interface {
	SystemMessage()
}

type Create struct {
	Failure error
}

func (p *Create) SystemMessage() {}

type Supervise struct {
	Child ActorRef
	Async bool
}

func (p *Supervise) SystemMessage() {}

type Failed struct {
	Child ActorRef
	Cause error
	Uid   int
}

func (p *Failed) SystemMessage() {}

type DeathWatchNotification struct {
	Actor              ActorRef
	ExistenceConfirmed bool
	addressTerminated  bool
}

func (p *DeathWatchNotification) SystemMessage() {}
