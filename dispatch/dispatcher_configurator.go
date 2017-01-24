package dispatch

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/configuration"
)

type DispatcherConfigurator struct {
	instance akka.MessageDispatcher

	config        *configuration.Config
	prerequisites *akka.DispatcherPrerequisites
}

func (p *DispatcherConfigurator) Config() *configuration.Config {
	return p.config
}

func (p *DispatcherConfigurator) DispatcherPrerequisites() *akka.DispatcherPrerequisites {
	return p.prerequisites
}

func NewDispatcherConfigurator(
	config *configuration.Config,
	prerequisites *akka.DispatcherPrerequisites) akka.MessageDispatcherConfigurator {

	deadlineTime := config.GetTimeDuration("throughput-deadline-time")

	configurator := &DispatcherConfigurator{
		config:        config,
		prerequisites: prerequisites,
	}

	instance := NewDispatcher(
		configurator,
		config.GetString("id"),
		int(config.GetInt64("throughput")),
		deadlineTime,
		NewThreadPoolConfig(10, 10),
	)

	configurator.instance = instance

	return configurator
}

func (p *DispatcherConfigurator) Dispatcher() akka.MessageDispatcher {
	return p.instance
}
