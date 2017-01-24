package dispatch

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/pkg/class_loader"
	"time"
)

func init() {
	class_loader.Default.Register((*Dispatcher)(nil), "dispatcher")
}

type Dispatcher struct {
	configurator            akka.MessageDispatcherConfigurator
	executorServiceDelegate *LazyExecutorServiceDelegate

	throughput             int
	throughputDeadlineTime time.Duration
}

func NewDispatcher(
	configurator akka.MessageDispatcherConfigurator,
	id string,
	throughput int,
	throughputDeadlineTime time.Duration,
	executorServiceFactoryProvider ExecutorServiceFactoryProvider,
) akka.MessageDispatcher {

	return &Dispatcher{
		configurator:            configurator,
		executorServiceDelegate: NewLazyExecutorServiceDelegate(executorServiceFactoryProvider.CreateExecutorServiceFactory(id)),
	}
}

func (p *Dispatcher) Attach(actor akka.ActorCell) {
	//TODO: register
	p.RegisterForExecution(actor.Mailbox(), false, true)
}

func (p *Dispatcher) Detach(actor akka.ActorCell) {
	return
}

func (p *Dispatcher) EventStream() akka.EventStream {
	return nil
}

func (p *Dispatcher) Execute(runnable akka.Runnable) {
	return
}

func (p *Dispatcher) Mailboxes() {
	return
}

func (p *Dispatcher) RegisterForExecution(mailbox akka.Mailbox, hasMessageHint bool, hasSystemMessageHint bool) bool {

	if mailbox.CanBeScheduledForExecution(hasMessageHint, hasSystemMessageHint) {
		if mailbox.SetAsScheduled() {
			p.executorService().Execute(mailbox)
			return true
		}
	}

	return false
}

func (p *Dispatcher) CreateMailbox(actor akka.Cell, mailboxType akka.MailboxType) akka.Mailbox {
	return newMailbox(mailboxType.Create(actor.Self(), actor.System()))
}

func (p *Dispatcher) Throughput() int {
	return p.throughput
}

func (p *Dispatcher) ThroughputTimeout() time.Duration {
	return p.throughputDeadlineTime
}

func (p *Dispatcher) Dispatch(receiver akka.ActorCell, invocation akka.Envelope) (err error) {
	mbox := receiver.Mailbox()
	if err = mbox.Enqueue(receiver.Self(), invocation); err != nil {
		return
	}

	p.RegisterForExecution(mbox, true, false)
	return
}

func (p *Dispatcher) SystemDispatch(receiver akka.ActorCell, invocation akka.SystemMessage) (err error) {
	mbox := receiver.Mailbox()
	if err = mbox.SystemEnqueue(receiver.Self(), invocation); err != nil {
		return
	}

	p.RegisterForExecution(mbox, false, true)
	return
}

func (p *Dispatcher) executorService() ExecutorServiceDelegate {
	return p.executorServiceDelegate
}
