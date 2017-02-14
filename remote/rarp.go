package remote

import (
	"github.com/go-akka/akka"
	"reflect"
)

var (
	_ akka.Extension   = (*RARP)(nil)
	_ akka.ExtensionId = (*RARP)(nil)
)

type RARP struct {
	provider akka.RemoteActorRefProvider
}

func NewRARP(provider akka.RemoteActorRefProvider) *RARP {
	return &RARP{
		provider: provider,
	}
}

func (p *RARP) ConfigureDispatcher(props akka.Props) akka.Props {
	return p.provider.RemoteSettings().ConfigureDispatcher(props)
}

func (p *RARP) Apply(system akka.ActorSystem) akka.Extension {
	return system.RegisterExtension(p)
}

func (p *RARP) Get(system akka.ActorSystem) akka.Extension {
	return p.Apply(system)
}

func (p *RARP) ExtensionType() reflect.Type {
	return reflect.TypeOf(p)
}

func (p *RARP) CreateExtension(system akka.ExtendedActorSystem) akka.Extension {
	return NewRARP(system.Provider().(akka.RemoteActorRefProvider))
}

func (p *RARP) Provider() akka.RemoteActorRefProvider {
	return p.provider
}

func (p *RARP) Lookup() akka.ExtensionId {
	return &RARP{}
}

func (p *RARP) Extension() {}
