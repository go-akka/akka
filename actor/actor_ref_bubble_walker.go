package actor

import (
	"fmt"
	"github.com/go-akka/akka"
)

type bubbleWalker struct {
	*akka.MinimalActorRef
}

func NewBubbleWalker(path akka.ActorPath, provider akka.ActorRefProvider) *bubbleWalker {
	base := akka.NewMinimalActorRef(path, provider)
	return &bubbleWalker{
		MinimalActorRef: base,
	}
}

func (p *bubbleWalker) IsWalking() bool {
	return true
}

func (p *bubbleWalker) Stop() {
	return
}

func (p *bubbleWalker) Tell(message interface{}, sender ...akka.ActorRef) (err error) {
	if p.IsWalking() {
		if message == nil {
			err = fmt.Errorf("Message is null")
			return
		}
	}
	return
}

func (p *bubbleWalker) sendSystemMessage(message akka.SystemMessage) {
	if p.IsWalking() {
		switch v := message.(type) {
		case *akka.Failed:
			{
				if internalActorRef, ok := v.Child.(akka.InternalActorRef); ok {
					internalActorRef.Stop()
				}
			}
		case *akka.Supervise:
			{

			}
		case *akka.DeathWatchNotification:
			{
				p.Stop()
			}
		default:
		}
	}
}
