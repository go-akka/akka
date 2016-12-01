package actor

import (
	"sync"
	"time"
)

var (
	_ ActorSystem = (*ActorSystemImpl)(nil)
)

type ActorSystemImpl struct {
	name        string
	startedTime time.Time

	settings *Settings
}

func NewActorSystem(name string, configs ...Config) (system ActorSystem, err error) {
	sys := &ActorSystemImpl{
		name:        name,
		startedTime: time.Now(),
	}

	var setting Settings
	if setting, err = NewSettings(nil, name); err != nil {
		return
	}

	sys.settings = &setting

	system = sys

	return
}

func (p *ActorSystemImpl) Settings() *Settings {
	return p.settings
}

func (p *ActorSystemImpl) ActorOf(props Props, name string) {
	return
}

func (p *ActorSystemImpl) ActorSelection(path ActorPath) {
	return
}

func (p *ActorSystemImpl) Child(child string) (path ActorPath) {
	return
}

func (p *ActorSystemImpl) Terminate() (wg sync.WaitGroup) {
	return
}

func (p *ActorSystemImpl) Name() string {
	return p.name
}

func (p *ActorSystemImpl) Stop() {
	return
}

func (p *ActorSystemImpl) Log() {
	return
}

func (p *ActorSystemImpl) DeadLetters() ActorRef {
	return nil
}

func (p *ActorSystemImpl) StartTime() int64 {
	return p.startedTime.Unix()
}

func (p *ActorSystemImpl) Uptime() time.Duration {
	return time.Now().Sub(p.startedTime)
}
