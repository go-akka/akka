package dispatch

import (
	"github.com/go-akka/concurrent"
	"sync"
	"time"
)

type ExecutorServiceDelegate interface {
	concurrent.ExecutorService
}

type LazyExecutorServiceDelegate struct {
	concurrent.ExecutorService

	factory  ExecutorServiceFactory
	initOnce sync.Once
}

func NewLazyExecutorServiceDelegate(factory ExecutorServiceFactory) *LazyExecutorServiceDelegate {
	return &LazyExecutorServiceDelegate{factory: factory}
}

func (p *LazyExecutorServiceDelegate) Copy() *LazyExecutorServiceDelegate {
	return &LazyExecutorServiceDelegate{factory: p.factory}
}

func (p *LazyExecutorServiceDelegate) Executor() concurrent.ExecutorService {
	p.initOnce.Do(func() {
		p.ExecutorService = p.factory.CreateExecutorService()
	})
	return p.ExecutorService
}

func (p *LazyExecutorServiceDelegate) Execute(command interface{}) {
	p.Executor().Execute(command)
}

func (p *LazyExecutorServiceDelegate) AwaitTermination(timeout time.Duration) (terminated bool) {
	return p.Executor().AwaitTermination(timeout)
}

func (p *LazyExecutorServiceDelegate) InvokeAll(tasks []interface{}) (future []concurrent.Future, err error) {
	return p.Executor().InvokeAll(tasks)
}

func (p *LazyExecutorServiceDelegate) InvokeAllDuration(tasks []interface{}, timeout time.Duration) (future []concurrent.Future, err error) {
	return p.Executor().InvokeAllDuration(tasks, timeout)
}

func (p *LazyExecutorServiceDelegate) InvokeAny(tasks []interface{}) (future []concurrent.Future, err error) {
	return p.Executor().InvokeAny(tasks)
}

func (p *LazyExecutorServiceDelegate) InvokeAnyDuration(tasks []interface{}, timeout time.Duration) (future []concurrent.Future, err error) {
	return p.Executor().InvokeAnyDuration(tasks, timeout)
}

func (p *LazyExecutorServiceDelegate) IsShutdown() bool {
	return p.Executor().IsShutdown()
}

func (p *LazyExecutorServiceDelegate) IsTerminated() bool {
	return p.Executor().IsTerminated()
}

func (p *LazyExecutorServiceDelegate) Shutdown() (err error) {
	return p.Executor().Shutdown()
}

func (p *LazyExecutorServiceDelegate) ShutdownNow() (runnables []concurrent.Runnable, err error) {
	return p.Executor().ShutdownNow()
}

func (p *LazyExecutorServiceDelegate) Submit(task interface{}) (future concurrent.Future, err error) {
	return p.Executor().Submit(task)
}
