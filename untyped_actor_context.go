package akka

type UntypedReceive func(message interface{})

type UntypedActorContext interface {
	ActorContext

	UntypedBecome(behavior UntypedReceive, discardOld bool)
}
