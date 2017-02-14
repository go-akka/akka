package remote

import (
	"github.com/go-akka/akka"
)

type RemotingCommand interface {
	RemotingCommand()
}

type Listen struct {
}

func (p *Listen) RemotingCommand() {}

type StartupFinished struct {
}

func (p *StartupFinished) RemotingCommand() {}

type ShutdownAndFlush struct{}

func (p *ShutdownAndFlush) RemotingCommand() {}

type Send struct {
	Message   interface{}
	Sender    akka.ActorRef
	Recipient RemoteActorRef
	seq       *SeqNo
}

func (p *Send) RemotingCommand() {}
func (p *Send) Seq() *SeqNo      { return p.seq }
func (p *Send) Copy(opt *SeqNo) *Send {
	return &Send{
		Message:   p.Message,
		Recipient: p.Recipient,
		Sender:    p.Sender,
		seq:       p.seq,
	}
}

type ManagementCommand struct {
	Cmd interface{}
}

func (p *ManagementCommand) RemotingCommand() {}

type ManagementCommandAck struct {
	Status bool
}
