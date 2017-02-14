package remote

import (
	"github.com/go-akka/akka"
)

type RemoteTransport interface {
	Provider() RemoteActorRefProvider

	System() akka.ExtendedActorSystem
	Start()

	Addresses() []akka.Address
	Send(message interface{}, sender akka.ActorRef, recipient RemoteActorRef)
	ManagementCommand(cmd interface{})
	LocalAddressForRemote(remote akka.Address) akka.Address
	DefaultAddress() akka.Address

	Log() akka.LoggingAdapter
	Quarantine(address akka.Address, uid int, reason string)
}
