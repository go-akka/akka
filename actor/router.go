package actor

type Routee interface {
	Send(message interface{}, sender ActorRef)
}

type RoutingLogic interface {
	Select(message interface{}, routees ...Routee) Routee
}

type Router struct {
}

func NewRouter(logic RoutingLogic, routees ...Routee) (router Router, err error) {
	return
}

func (p Router) AddRoutee(sel ActorSelection) {
}
