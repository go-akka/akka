package remote

import (
	"github.com/go-akka/akka"
)

type RemoteActorRef interface {
	akka.InternalActorRef
}

type RemoteActorRefProvider interface {
	akka.ActorRefProvider
	RemoteSettings() *RemoteSettings
	RemoteActorRefProvider()
}
