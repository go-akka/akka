package akka

type Routee interface {
	Send(message interface{}, sender ActorRef)
}

type RoutingLogic interface {
	Select(message interface{}, routees ...Routee) Routee
}
