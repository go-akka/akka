package dispatch

import (
	"github.com/go-akka/concurrent"
)

type ExecutorServiceFactory interface {
	Produce(id string) concurrent.ExecutorService
}
