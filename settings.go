package akka

import (
	"github.com/go-akka/configuration"
)

type Settings struct {
	// system ActorSystem
	name   string
	config *configuration.Config

	userConfig     *configuration.Config
	fallbackConfig *configuration.Config

	ConfigVersion           string
	ProviderClass           string
	SupervisorStrategyClass string
	DebugEventStream        bool
	LogLevel                string
	SchedulerClass          string

	LoggersDispatcher string

	Loggers []string
}

func NewSettings(systemName string, config *configuration.Config) (settings *Settings, err error) {
	s := &Settings{
		userConfig: config,
		// system:           system,
		name:             systemName,
		DebugEventStream: config.GetBoolean("akka.actor.debug.event-stream", false),
		ProviderClass:    config.GetString("akka.actor.provider"),
		LogLevel:         config.GetString("akka.loglevel"),
		SchedulerClass:   config.GetString("akka.scheduler.implementation"),
	}
	settings = s
	return
}

func (p *Settings) rebuildConfig() {
	p.config = p.userConfig.WithFallback(p.fallbackConfig)
}

func (p *Settings) String() string {
	return p.config.String()
}

func (p *Settings) Config() *configuration.Config {
	return p.config
}
