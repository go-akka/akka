package dispatch

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/pkg/dynamic_access"
	"github.com/go-akka/configuration"
	"github.com/orcaman/concurrent-map"
	"strings"
)

const (
	DefaultDispatcherId = "akka.actor.default-dispatcher"
)

func NewDefaultDispatcherPrerequisites(
	eventStream akka.EventStream,
	scheduler akka.Scheduler,
	dynamicAccess dynamic_access.DynamicAccess,
	settings *akka.Settings,
	mailboxes akka.Mailboxes,
) *akka.DispatcherPrerequisites {

	return &akka.DispatcherPrerequisites{
		EventStream:   eventStream,
		Scheduler:     scheduler,
		DynamicAccess: dynamicAccess,
		Settings:      settings,
		Mailboxes:     mailboxes,
	}
}

type Dispatchers struct {
	settings      *akka.Settings
	prerequisites *akka.DispatcherPrerequisites

	defaultDispatcherConfig *configuration.Config
	dispatcherConfigurators cmap.ConcurrentMap
}

func NewDispatchers(settings *akka.Settings, prerequisites *akka.DispatcherPrerequisites) akka.Dispatchers {
	dispatcher := &Dispatchers{
		settings:                settings,
		prerequisites:           prerequisites,
		dispatcherConfigurators: cmap.New(),
	}

	dispatcher.defaultDispatcherConfig =
		dispatcher.idConfig(DefaultDispatcherId).
			WithFallback(dispatcher.settings.Config().GetConfig(DefaultDispatcherId))

	return dispatcher
}

func (p *Dispatchers) Lookup(id string) akka.MessageDispatcher {
	return p.lookupConfigurator(id).Dispatcher()
}

func (p *Dispatchers) HasDispatcher(id string) bool {
	return true
}

func (p *Dispatchers) RegisterConfigurator(id string, configurator akka.MessageDispatcherConfigurator) bool {
	return false
}

func (p *Dispatchers) defaultGlobalDispatcher() akka.MessageDispatcher {
	return p.Lookup(DefaultDispatcherId)
}

func (p *Dispatchers) lookupConfigurator(id string) akka.MessageDispatcherConfigurator {
	configurator, exist := p.dispatcherConfigurators.Get(id)

	if !exist {

		newConfigurator := p.configuratorFrom(p.config(id, p.settings.Config().GetConfig(id)))

		p.dispatcherConfigurators.SetIfAbsent(id, newConfigurator)

		return newConfigurator
	}
	return configurator.(akka.MessageDispatcherConfigurator)
}

func (p *Dispatchers) idConfig(id string) *configuration.Config {

	return configuration.ParseString("id:" + id)
}

func (p *Dispatchers) config(id string, appConfig *configuration.Config) *configuration.Config {
	simpleName := string([]byte(id)[strings.LastIndex(id, ".")+1:])

	return p.idConfig(id).
		WithFallback(appConfig).
		WithFallback(configuration.ParseString("name:" + simpleName)).
		WithFallback(p.defaultDispatcherConfig)
}

func (p *Dispatchers) configuratorFrom(cfg *configuration.Config) akka.MessageDispatcherConfigurator {
	if !cfg.HasPath("id") {
		panic("Missing dispatcher 'id' property in config: " + cfg.Root().String())
	}

	switch cfg.GetString("type") {
	case "dispatcher":
		{
			return NewDispatcherConfigurator(cfg, p.prerequisites)
		}
	}

	return nil
}
