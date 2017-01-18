package dispatch

import (
	"github.com/go-akka/concurrent"
)

type ExecutorServiceDelegate struct {
	concurrent.ExecutorService
}

func (p *ExecutorServiceDelegate) Executor() concurrent.ExecutorService {
	return p.ExecutorService
}
