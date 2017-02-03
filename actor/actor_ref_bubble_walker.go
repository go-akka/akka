package actor

import (
	"fmt"

	"github.com/go-akka/akka"
	"github.com/go-akka/akka/dispatch/sysmsg"
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
		case *sysmsg.Failed:
			{
				if internalActorRef, ok := v.Child.(akka.InternalActorRef); ok {
					internalActorRef.Stop()
				}
			}
		case *sysmsg.Supervise:
			{

			}
		case *sysmsg.DeathWatchNotification:
			{
				p.Stop()
			}
		default:
		}
	}
}
