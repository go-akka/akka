package akka

import (
	"time"
)

type MessageDispatcher interface {
	Attach(actor Cell)
	Detach(actor Cell)
	EventStream() EventStream
	Execute(runnable Runnable)
	Mailboxes()
	RegisterForExecution(mailbox Mailbox, hasMessageHint bool, hasSystemMessageHint bool) bool

	Throughput() int
	ThroughputTimeout() time.Duration
}
