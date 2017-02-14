package remote

type DisassociateInfo int

const (
	UnknownDisassociateInfo     DisassociateInfo = 0
	ShutdownDisassociateInfo    DisassociateInfo = 1
	QuarantinedDisassociateInfo DisassociateInfo = 2
)

type HandleEvent interface {
	HandleEvent()
}

type AssociationHandle struct {
}

type InboundAssociation struct {
	Association *AssociationHandle
}

func (p *InboundAssociation) AssociationEvent() {}

type Disassociated struct {
	Info DisassociateInfo
}

func (p *Disassociated) HandleEvent() {}

type UnderlyingTransportError struct {
	Cause   error
	Message string
}

func (p *UnderlyingTransportError) HandleEvent() {}
