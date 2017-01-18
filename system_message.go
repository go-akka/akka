package akka

type Terminated struct {
	actor              ActorRef
	addressTerminated  bool
	existenceConfirmed bool
}

func (p *Terminated) AutoReceivedMessage() {
}

// Actor is the watched actor that terminated
func (p *Terminated) Actor() ActorRef {
	return p.actor
}

// AddressTerminated is about the Terminated message was derived from that the remote node hosting the watched actor was detected as unreachable
func (p *Terminated) AddressTerminated() bool {
	return p.addressTerminated
}

// GetExistenceConfirmed is false when the Terminated message was not sent directly from the watched actor,
// but derived from another source, such as when watching a non-local ActorRef,
// which might not have been resolved
func (p *Terminated) ExistenceConfirmed() bool {
	return p.existenceConfirmed
}

type StopChild struct {
	child ActorRef
}

func (p *StopChild) Child() ActorRef {
	return p.child
}

type AddressTerminated struct {
}

func (p *AddressTerminated) AutoReceivedMessage() {
}

type Kill struct {
}

func (p *Kill) AutoReceivedMessage() {
}

type PoisonPill struct {
}

func (p *PoisonPill) AutoReceivedMessage() {
}

type ActorSelectionMessage struct {
}

func (p *ActorSelectionMessage) AutoReceivedMessage() {
}

type Identify struct {
}

func (p *Identify) AutoReceivedMessage() {
}
