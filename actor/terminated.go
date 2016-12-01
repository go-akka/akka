package actor

type Terminated struct {
	actor              ActorRef
	addressTerminated  bool
	existenceConfirmed bool
}

// Actor is the watched actor that terminated
func (p *Terminated) GetActor() ActorRef {
	return p.actor
}

// AddressTerminated is about the Terminated message was derived from that the remote node hosting the watched actor was detected as unreachable
func (p *Terminated) GetAddressTerminated() bool {
	return p.addressTerminated
}

// GetExistenceConfirmed is false when the Terminated message was not sent directly from the watched actor,
// but derived from another source, such as when watching a non-local ActorRef,
// which might not have been resolved
func (p *Terminated) GetExistenceConfirmed() bool {
	return p.existenceConfirmed
}
