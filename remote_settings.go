package akka

import (
	"fmt"
	"github.com/go-akka/configuration"
	"strings"
)

type TransportSettings struct {
	TransportClass string
	Adapters       []string
	Config         *configuration.Config
}

func NewTransportSettings(config *configuration.Config) *TransportSettings {
	settings := &TransportSettings{
		TransportClass: config.GetString("transport-class"),
		Config:         config,
	}

	settings.Adapters = reverseStringList(config.GetStringList("applied-adapters"))

	return settings
}

type RemoteSettings struct {
	config *configuration.Config

	Dispatcher                    string
	ProviderClass                 string
	TransportNames                []string
	Transports                    []*TransportSettings
	Adapters                      map[string]string
	RemoteLifecycleEventsLogLevel string
}

func NewRemoteSettings(config *configuration.Config) (settings *RemoteSettings, err error) {
	s := &RemoteSettings{
		config: config,
	}

	s.ProviderClass = config.GetString("akka.remote.use-dispatcher")
	s.TransportNames = config.GetStringList("akka.remote.enabled-transports")
	s.RemoteLifecycleEventsLogLevel = config.GetString("akka.remote.log-remote-lifecycle-events", "DEBUG")
	s.Dispatcher = config.GetString("akka.remote.use-dispatcher")
	if strings.ToUpper(s.RemoteLifecycleEventsLogLevel) == "ON" {
		s.RemoteLifecycleEventsLogLevel = "DEBUG"
	}

	for i := 0; i < len(s.TransportNames); i++ {
		transportConfig := s.transportConfigFor(s.TransportNames[i])
		s.Transports = append(s.Transports, NewTransportSettings(transportConfig))
	}

	s.Adapters = s.configToMap(config.GetConfig("akka.remote.adapters"))

	settings = s

	return
}

func (p *RemoteSettings) String() string {
	return p.config.String()
}

func (p *RemoteSettings) Config() *configuration.Config {
	return p.config
}

func (p *RemoteSettings) ConfigureDispatcher(props Props) Props {
	if len(p.Dispatcher) == 0 {
		return props
	}
	return props.WithDispatcher(p.Dispatcher)
}

func (p *RemoteSettings) transportConfigFor(transportName string) *configuration.Config {
	return p.config.GetConfig(transportName)
}

func (p *RemoteSettings) configToMap(config *configuration.Config) map[string]string {
	if config == nil || config.IsEmpty() {
		return map[string]string{}
	}

	unwrapped := config.Root().GetObject().Unwrapped()
	retVal := map[string]string{}
	for k, v := range unwrapped {
		if v != nil {
			retVal[k] = fmt.Sprintf("%s", v)
		} else {
			retVal[k] = ""
		}
	}

	return retVal
}
