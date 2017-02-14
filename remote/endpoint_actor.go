package remote

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/actor"
)

type StopReading struct {
	Writer  akka.ActorRef
	ReplyTo akka.ActorRef
}

type EndpointActor struct {
	*actor.UntypedActor

	LocalAddress  akka.Address
	RemoteAddress akka.Address
	Settings      *akka.RemoteSettings
	// Transport     AkkaProtocolTransport
	log            akka.LoggingAdapter
	EventPublisher EventPublisher
	Inbound        bool
}

func NewEndpointActor(localAddress akka.Address, remoteAddress akka.Address, settings *akka.RemoteSettings) *EndpointActor {
	return &EndpointActor{
		LocalAddress:  localAddress,
		RemoteAddress: remoteAddress,
		Settings:      settings,
	}
}

func (p *EndpointActor) publishError(ex error, level akka.LogLevel) {
	p.tryPublish(NewAssociationErrorEvent(ex, p.LocalAddress, p.RemoteAddress, p.Inbound, level))
}

func (p *EndpointActor) publishDisassociated() {
	p.tryPublish(NewDisassociatedEvent(p.LocalAddress, p.RemoteAddress, p.Inbound))
}

func (p *EndpointActor) tryPublish(ev RemotingLifecycleEvent) {
	if err := p.EventPublisher.NotifyListeners(ev); err != nil {
		p.log.Error(err, "Unable to publish error event to EventStream")
	}
}

type EndpointWriter struct {
	*EndpointActor

	reader akka.ActorRef
}

func (p *EndpointWriter) PreStart() (err error) {

	return
}

func (p *EndpointWriter) Receive(message interface{}) (handled bool, err error) {
	switch message.(type) {
	case *Send:
		{

		}
	}
	return
}

func (p *EndpointWriter) WritingReceive(message interface{}) (handled bool, err error) {
	switch message.(type) {
	case *Send:
		{

		}
	}
	return
}

func (p *EndpointWriter) BufferingReceive(message interface{}) (handled bool, err error) {
	switch message.(type) {
	case *Send:
		{

		}
	}
	return
}

// func (p *EndpointWriter) startReadEndpoint(handle AkkaProtocolHandle) {
// }

func (p *EndpointWriter) sendBufferedMessages() {
}

func (p *EndpointWriter) becomeWritingOrSendBufferedMessages() {
}

func (p *EndpointWriter) enqueueInBuffer(message interface{}) {
}

func (p *EndpointWriter) writeSend(send *Send) {
}

type EndpointReader struct {
	*EndpointActor
}

func (p *EndpointReader) PreStart() (err error) {

	return
}

func (e *EndpointReader) PostStop(err error) {
	return
}

func (p *EndpointReader) Receive(message interface{}) (handled bool, err error) {
	switch message.(type) {
	case *Disassociated:
		{

		}
	case *StopReading:
		{
			p.Become(p.NotReadingReceive, true)
		}
	}

	return
}

func (p *EndpointReader) NotReadingReceive(message interface{}) (handled bool, err error) {
	switch message.(type) {
	case *Disassociated:
		{

		}
	case *StopReading:
		{

		}
	}

	return
}
