package sysmsg

import (
	"github.com/go-akka/akka"
	"strconv"
)

type NoMessage struct{}

func (p *NoMessage) SystemMessage() {}
func (p *NoMessage) String() string {
	return "NoMessage"
}

type Create struct {
	Failure error
}

func (p *Create) SystemMessage() {}

type Supervise struct {
	Child akka.ActorRef
	Async bool
}

func (p *Supervise) SystemMessage() {}
func (p *Supervise) String() string {

	strAsync := "False"
	if p.Async {
		strAsync = "True"
	}

	return "<Supervise>: " + p.Child.Path().String() + ", Async=" + strAsync
}

type Failed struct {
	Child akka.ActorRef
	Cause error
	Uid   int
}

func (p *Failed) SystemMessage() {}
func (p *Failed) String() string {
	str := "<Failed>: " + p.Child.Path().String() + " (" + strconv.Itoa(p.Uid) + ") "
	if p.Cause != nil {
		str += ", Cause=" + p.Cause.Error()
	}
	return str
}

type DeathWatchNotification struct {
	Actor              akka.ActorRef
	ExistenceConfirmed bool
	addressTerminated  bool
}

func (p *DeathWatchNotification) SystemMessage() {}

type Stop struct{}

func (p *Stop) SystemMessage() {}
func (p *Stop) String() string {
	return "<Stop>"
}

type StopChild struct {
	child akka.ActorRef
}

func (p *StopChild) Child() akka.ActorRef {
	return p.child
}
func (p *StopChild) SystemMessage() {}
func (p *StopChild) String() string {
	return "<StopChild> " + p.child.Path().String()
}

type Terminate struct{}

func (p *Terminate) SystemMessage() {}
func (p *Terminate) String() string {
	return "<ActorSelectionMessage>"
}
