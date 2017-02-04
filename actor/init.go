package actor

import (
	"github.com/go-akka/akka/actor/props"
	"github.com/go-akka/akka/pkg/class_loader"
)

func init() {
	class_loader.Default.Register((*LocalActorRefProvider)(nil), "LocalActorRefProvider")
	props.RegisterGlobalProducerCreator(newReflectProducer)
}
