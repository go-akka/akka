package remote

import (
	"fmt"
	"github.com/go-akka/akka"
)

type RemotingLifecycleEvent interface {
	LogLevel() akka.LogLevel
}

type AssociationEvent struct {
	localAddress  akka.Address
	remoteAddress akka.Address
	inbound       bool
	eventName     string
}

func (p *AssociationEvent) String() string {
	dirChar := "->"
	if p.inbound {
		dirChar = "<-"
	}
	return fmt.Sprintf("%s [%s] %s %s", p.eventName, p.localAddress.String(), dirChar, p.remoteAddress.String())
}

type AssociatedEvent struct {
	*AssociationEvent
}

func (p *AssociatedEvent) LogLevel() akka.LogLevel {
	return akka.DebugLevel
}

func (p *AssociatedEvent) LocalAddress() akka.Address {
	return p.localAddress
}

func (p *AssociatedEvent) RemoteAddress() akka.Address {
	return p.remoteAddress
}

func NewAssociatedEvent(localAddress akka.Address, remoteAddress akka.Address, inbound bool) *AssociatedEvent {
	return &AssociatedEvent{
		AssociationEvent: &AssociationEvent{
			localAddress:  localAddress,
			remoteAddress: remoteAddress,
			inbound:       inbound,
			eventName:     "Associated",
		},
	}
}

type DisassociatedEvent struct {
	*AssociationEvent
}

func (p *DisassociatedEvent) LogLevel() akka.LogLevel {
	return akka.DebugLevel
}

func (p *DisassociatedEvent) LocalAddress() akka.Address {
	return p.localAddress
}

func (p *DisassociatedEvent) RemoteAddress() akka.Address {
	return p.remoteAddress
}

func NewDisassociatedEvent(localAddress akka.Address, remoteAddress akka.Address, inbound bool) *DisassociatedEvent {
	return &DisassociatedEvent{
		AssociationEvent: &AssociationEvent{
			localAddress:  localAddress,
			remoteAddress: remoteAddress,
			inbound:       inbound,
			eventName:     "Disassociated",
		},
	}
}

type AssociationErrorEvent struct {
	*AssociationEvent
	level akka.LogLevel
	cause error
}

func (p *AssociationErrorEvent) LogLevel() akka.LogLevel {
	return p.level
}

func (p *AssociationErrorEvent) LocalAddress() akka.Address {
	return p.localAddress
}

func (p *AssociationErrorEvent) RemoteAddress() akka.Address {
	return p.remoteAddress
}

func (p *AssociationErrorEvent) Cause() error {
	return p.cause
}

func (p *AssociationErrorEvent) String() string {
	// TODO: add error stacktrace
	return fmt.Sprintf("%s: Error %s", p.AssociationEvent.String(), p.cause.Error())
}

func NewAssociationErrorEvent(cause error, localAddress akka.Address, remoteAddress akka.Address, inbound bool, level akka.LogLevel) *AssociationErrorEvent {
	return &AssociationErrorEvent{
		AssociationEvent: &AssociationEvent{
			localAddress:  localAddress,
			remoteAddress: remoteAddress,
			inbound:       inbound,
			eventName:     "AssociationError",
		},
		level: level,
		cause: cause,
	}
}

type EventPublisher struct {
	system   akka.ActorSystem
	log      akka.LoggingAdapter
	logLevel akka.LogLevel
}

func NewEventPlublisher(system akka.ActorSystem, log akka.LoggingAdapter, logLevel akka.LogLevel) *EventPublisher {
	return &EventPublisher{
		system:   system,
		log:      log,
		logLevel: logLevel,
	}
}

func (p *EventPublisher) NotifyListeners(message RemotingLifecycleEvent) (err error) {
	p.system.EventStream().Publish(message)
	if message.LogLevel() >= p.logLevel {
		p.log.Log(message.LogLevel(), "%s", message)
	}

	return
}
