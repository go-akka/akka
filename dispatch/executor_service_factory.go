package dispatch

import (
	"github.com/go-akka/concurrent"
)

type ExecutorServiceFactory interface {
	CreateExecutorService() concurrent.ExecutorService
}

type ExecutorServiceFactoryProvider interface {
	CreateExecutorServiceFactory(id string) ExecutorServiceFactory
}
