package akka

import (
	"github.com/go-akka/configuration"
	"time"
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
	LogLevel                string
	SchedulerClass          string
	StdoutLogLevel          string
	LoggerStartTimeout      time.Duration

	DebugUnhandledMessage bool
	DebugEventStream      bool
	DebugAutoReceive      bool
	DebugLifecycle        bool

	LoggersDispatcher string

	Loggers []string
}

func NewSettings(systemName string, config *configuration.Config) (settings *Settings, err error) {
	s := &Settings{
		userConfig: config,
		name:       systemName,
	}

	s.rebuildConfig()

	s.ProviderClass = config.GetString("akka.actor.provider")
	s.LogLevel = config.GetString("akka.loglevel")
	s.SchedulerClass = config.GetString("akka.scheduler.implementation")

	s.LogLevel = config.GetString("akka.loglevel")
	s.StdoutLogLevel = config.GetString("akka.stdout-loglevel")
	s.Loggers = config.GetStringList("akka.loggers")
	s.LoggersDispatcher = config.GetString("akka.loggers-dispatcher")
	s.LoggerStartTimeout = config.GetTimeDuration("akka.logger-startup-timeout")

	s.DebugEventStream = config.GetBoolean("akka.actor.debug.event-stream", false)
	s.DebugUnhandledMessage = config.GetBoolean("akka.actor.debug.unhandled")
	s.DebugAutoReceive = config.GetBoolean("akka.actor.debug.autoreceive")
	s.DebugLifecycle = config.GetBoolean("akka.actor.debug.lifecycle")

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
