package routing

import (
	. "github.com/go-akka/akka"
)

type NoRoutee struct {
}

func (p NoRoutee) Send(message interface{}, sender ActorRef) {
}

type Router struct {
}

func NewRouter(logic RoutingLogic, routees ...Routee) (router Router, err error) {
	return
}

func (p Router) AddRoutee(sel ActorSelection) {
}
