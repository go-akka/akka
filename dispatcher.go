package akka

import (
	"github.com/go-akka/akka/pkg/dynamic_access"
	"github.com/go-akka/configuration"
)

type DispatcherPrerequisites struct {
	EventStream   EventStream
	Scheduler     Scheduler
	DynamicAccess dynamic_access.DynamicAccess
	Settings      *Settings
	Mailboxes     Mailboxes
}

type MessageDispatcherConfigurator interface {
	Config() *configuration.Config
	DispatcherPrerequisites() *DispatcherPrerequisites
	Dispatcher() MessageDispatcher
}
