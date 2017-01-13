package akka

import (
	"github.com/go-akka/akka/pkg/dynamic_access"
)

type Mailboxes struct {
	setting     Settings
	eventStream EventStream
}

func NewMailboxes(
	settings Settings,
	eventStream EventStream,
	dynamicAccess dynamic_access.DynamicAccess,
	deadLetters ActorRef) {
}
