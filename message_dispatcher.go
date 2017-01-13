package akka

import (
	"github.com/go-akka/configuration"
)

var (
	_ DispatcherPrerequisites = (*DefaultDispatcherPrerequisites)(nil)
)

const (
	DefaultThroughput = 100
)

type DispatcherPrerequisites interface {
	EventStream() EventStream
	Mailboxes() Mailboxes
	Scheduler() Scheduler
	Settings() Settings
}

type DefaultDispatcherPrerequisites struct {
	mailboxes   Mailboxes
	settings    Settings
	scheduler   Scheduler
	eventStream EventStream
}

func (p *DefaultDispatcherPrerequisites) EventStream() EventStream {
	return p.eventStream
}

func (p *DefaultDispatcherPrerequisites) Mailboxes() Mailboxes {
	return p.mailboxes
}

func (p *DefaultDispatcherPrerequisites) Scheduler() Scheduler {
	return p.scheduler
}

func (p *DefaultDispatcherPrerequisites) Settings() Settings {
	return p.settings
}

type MessageDispatcherConfigurator struct {
	prerequisites DispatcherPrerequisites
	config        *configuration.Config
}

func NewMessageDispatcherConfigurator(config *configuration.Config, prerequisites DispatcherPrerequisites) *MessageDispatcherConfigurator {
	return &MessageDispatcherConfigurator{
		prerequisites: prerequisites,
		config:        config,
	}
}

type MessageDispatcher struct {
	configurator MessageDispatcherConfigurator
	Throughput   int
}

func NewMessageDispatcher(configurator MessageDispatcherConfigurator) *MessageDispatcher {
	return &MessageDispatcher{
		configurator: configurator,
		Throughput:   DefaultThroughput,
	}
}

func (p *MessageDispatcher) MessageDispatcherConfigurator() MessageDispatcherConfigurator {
	return p.configurator
}

func (p *MessageDispatcher) Dispatch(cell Cell, envelope Envelope) {
	return
}

func (p *MessageDispatcher) Attach(actor Cell) {

}
