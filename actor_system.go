package akka

import (
	"github.com/go-akka/akka/pkg/dynamic_access"
	"sync"
)

type ActorSystem interface {
	ActorRefFactory

	Settings() *Settings
	Name() string
	Log()
	DeadLetters() ActorRef
	Terminate() sync.WaitGroup

	EventStream() EventStream

	RegisterOnTermination(fn func())

	// Child is Create a new child actor path.
	Child(child string) (path ActorPath, err error)

	// Recursively create a descendant’s path by appending all child names.
	Descendant(names ...string) (path ActorPath, err error)

	// Start-up time in milliseconds.
	StartTime() int64

	// Up-time of this actor system in seconds.
	Uptime() int64

	RegisterExtension(ext ExtensionId) Extension
	Extension(ext ExtensionId) Extension
	HasExtension(ext ExtensionId) bool
}

type ActorSystemImpl interface {
	ExtendedActorSystem
}

type ExtendedActorSystem interface {
	ActorSystem

	Provider() ActorRefProvider
	Guardian() LocalActorRef
	SystemGuardian() LocalActorRef

	SystemActorOf(props Props, name string) (ref ActorRef, err error)
	DynamicAccess() dynamic_access.DynamicAccess

	LogFilter() LoggingFilter
}
