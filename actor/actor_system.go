package actor

import (
	"sync"
	"time"
)

type ActorSystem interface {
	Settings() *Settings
	ActorOf(props Props, name string)
	ActorSelection(path ActorPath)
	Name() string
	Stop()
	Log()
	DeadLetters() ActorRef
	Child(child string) ActorPath
	Terminate() sync.WaitGroup
	StartTime() int64
	Uptime() time.Duration
}
