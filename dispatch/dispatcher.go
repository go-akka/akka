package dispatch

import (
	"github.com/go-akka/akka"
)

type Dispatcher struct {
	configurator akka.MessageDispatcherConfigurator
}
