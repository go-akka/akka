package akka

type RouterConfig interface {
	CreateRouter(system ActorSystem)
	RouterDispatcher() string
	IsManagementMessage(msg interface{}) bool
	RoutingLogicController(routingLogic RoutingLogic) Props
	StopRouterWhenAllRouteesRemoved() bool
	VerifyConfig(path ActorPath) (err error)
	WithFallback(other RouterConfig) RouterConfig
}

type NoRouter struct {
}

func (p NoRouter) CreateRouter(system ActorSystem) {
	return
}

func (p NoRouter) RouterDispatcher() string {
	return ""
}

func (p NoRouter) IsManagementMessage(msg interface{}) bool {
	return false
}

func (p NoRouter) RoutingLogicController(routingLogic RoutingLogic) Props {
	return nil
}

func (p NoRouter) StopRouterWhenAllRouteesRemoved() bool {
	return false
}

func (p NoRouter) VerifyConfig(path ActorPath) (err error) {
	return
}

func (p NoRouter) WithFallback(other RouterConfig) RouterConfig {
	return nil
}
