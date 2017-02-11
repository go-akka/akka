package akka

import (
	"github.com/go-akka/configuration"
)

type RemoteTransport interface {
	Provider() RemoteActorRefProvider

	Config() configuration.Config
	System() ExtendedActorSystem
	Start()

	Addresses() []Address
	Send(message interface{}, sender ActorRef, recipient RemoteActorRef)
	ManagementCommand(cmd interface{})
	LocalAddressForRemote(remote Address) Address
	DefaultAddress() Address

	Log() LoggingAdapter
	Quarantine(address Address, uid int, reason string)
}
