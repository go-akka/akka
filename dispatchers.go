package akka

type Dispatchers interface {
	Lookup(id string) MessageDispatcher
	HasDispatcher(id string) bool
	RegisterConfigurator(id string, configurator MessageDispatcherConfigurator) bool
}
