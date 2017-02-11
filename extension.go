package akka

import (
	"reflect"
)

type Extension interface {
	Extension()
}

type ExtensionId interface {
	Apply(system ActorSystem) Extension
	Get(system ActorSystem) Extension
	CreateExtension(system ExtendedActorSystem) Extension
	ExtensionType() reflect.Type
}

type ExtensionProvider interface {
	Lookup() ExtensionId
}
