package remote

import (
	"github.com/go-akka/akka"
	"time"
)

type readOnlyAddress struct {
	Endpoint akka.ActorRef
	Uid      int
}

type EndpointRegistry struct {
	addressToWritable map[akka.Address]EndpointPolicy
	writableToAddress map[akka.ActorRef]akka.Address
	addressToReadonly map[akka.Address]readOnlyAddress
	readonlyToAddress map[akka.ActorRef]akka.Address
}

func (p *EndpointRegistry) RegisterWritableEndpoint(address akka.Address, endpoint akka.ActorRef, uid int, refuseUid int) akka.ActorRef {
	return nil
}

func (p *EndpointRegistry) RegisterWritableEndpointUid(remoteAddress akka.Address, uid int) {
	return
}

func (p *EndpointRegistry) RegisterWritableEndpointRefuseUid(remoteAddress akka.Address, refuseUid int) {
	return
}

func (p *EndpointRegistry) RegisterReadOnlyEndpoint(address akka.Address, endpoint akka.ActorRef, uid int) akka.ActorRef {
	return nil
}

func (p *EndpointRegistry) UnregisterEndpoint(endpoint akka.ActorRef) {
	return
}

func (p *EndpointRegistry) AddressForWriter(writer akka.ActorRef) {
	return
}

func (p *EndpointRegistry) ReadOnlyEndpointFor(address akka.Address) (endpoint akka.ActorRef, uid int) {
	return
}

func (p *EndpointRegistry) IsWritable(endpoint akka.ActorRef) bool {
	return false
}

func (p *EndpointRegistry) IsReadOnly(endpoint akka.ActorRef) bool {
	return false
}

func (p *EndpointRegistry) IsQuarantined(address akka.Address, uid int) bool {
	return false
}

func (p *EndpointRegistry) RefuseUid(address akka.Address) int {
	return 0
}

func (p *EndpointRegistry) MarkAsFailed(endpoint akka.ActorRef, timeOfRelease time.Time) {
	return
}

func (p *EndpointRegistry) MarkAsQuarantined(address akka.Address, uid int, timeOfRelease time.Time) {
	return
}

func (p *EndpointRegistry) RemovePolicy(address akka.Address) {
	return
}

func (p *EndpointRegistry) AllEndpoints() []akka.ActorRef {
	return nil
}

func (p *EndpointRegistry) Prune() {
	return
}
