package dispatch

import (
	"github.com/go-akka/concurrent"
)

var (
	_ ExecutorServiceFactoryProvider = (*ThreadPoolConfig)(nil)
)

type ThreadPoolConfig struct {
	nRoutine  int
	queueSize int
}

type ThreadPoolExecutorServiceFactory struct {
	nRoutine  int
	queueSize int
}

func (p *ThreadPoolExecutorServiceFactory) CreateExecutorService() concurrent.ExecutorService {
	return concurrent.NewFixedRoutinePool(p.nRoutine, p.queueSize)
}

func NewThreadPoolConfig(nRoutine, queueSize int) ExecutorServiceFactoryProvider {
	return &ThreadPoolConfig{
		nRoutine:  nRoutine,
		queueSize: queueSize,
	}
}

func (p *ThreadPoolConfig) CreateExecutorServiceFactory(id string) ExecutorServiceFactory {
	return &ThreadPoolExecutorServiceFactory{
		nRoutine:  p.nRoutine,
		queueSize: p.queueSize,
	}
}
