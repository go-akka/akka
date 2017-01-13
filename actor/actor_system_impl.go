package actor

import (
	"regexp"
	"sync"
	"time"

	"github.com/go-akka/akka"
	"github.com/go-akka/akka/pkg/class_loader"
	"github.com/go-akka/akka/pkg/dynamic_access"
	"github.com/go-akka/configuration"
)

type ActorSystemImpl struct {
	name        string
	startedTime time.Time

	settings *akka.Settings

	path akka.ActorPath

	classLoader   class_loader.ClassLoader
	dynamicAccess dynamic_access.DynamicAccess
	eventStream   akka.EventStream
	scheduler     akka.Scheduler

	provider       akka.ActorRefProvider
	lookupRoot     akka.InternalActorRef
	guardian       *LocalActorRef
	systemGuardian *LocalActorRef
}

func AkkaClassLoader() class_loader.ClassLoader {
	return class_loader.Default()
}

func NewActorSystem(name string, config ...*configuration.Config) (system *ActorSystemImpl, err error) {
	if match, _ := regexp.MatchString("^[a-zA-Z0-9][a-zA-Z0-9-_]*$", name); !match {
		err = akka.ErrInvalidActorSystemName
		return
	}

	classLoader := class_loader.NewClassicClassLoader(class_loader.Default())

	sys := &ActorSystemImpl{
		name:          name,
		startedTime:   time.Now(),
		classLoader:   classLoader,
		dynamicAccess: dynamic_access.NewReflectiveDynamicAccess(classLoader),
	}

	var conf *configuration.Config
	if len(config) > 0 {
		conf = config[0]
	}

	if err = sys.configureSettings(conf); err != nil {
		return
	}

	if err = sys.configureEventStream(); err != nil {
		return
	}
	// sys.configureLoggers()
	// if err = sys.configureScheduler(); err != nil {
	// 	return
	// }
	sys.configureProvider()
	// sys.configureMailboxes()
	// sys.configureDispatchers()

	system = sys

	return
}

func (p *ActorSystemImpl) Settings() *akka.Settings {
	return p.settings
}

func (p *ActorSystemImpl) Child(child string) (path akka.ActorPath, err error) {
	return
}

func (p *ActorSystemImpl) Descendant(names ...string) (path akka.ActorPath, err error) {
	return
}

func (p *ActorSystemImpl) Terminate() (wg sync.WaitGroup) {
	return
}

func (p *ActorSystemImpl) RegisterOnTermination(fn func()) {
	return
}

func (p *ActorSystemImpl) Name() string {
	return p.name
}

func (p *ActorSystemImpl) Log() {
	return
}

func (p *ActorSystemImpl) DeadLetters() akka.ActorRef {
	return nil
}

func (p *ActorSystemImpl) StartTime() int64 {
	return p.startedTime.Unix()
}

func (p *ActorSystemImpl) Uptime() int64 {
	return int64(time.Now().Sub(p.startedTime).Seconds())
}

func (p *ActorSystemImpl) Forward(message interface{}) {
	return
}

func (p *ActorSystemImpl) Equals(that interface{}) bool {
	switch other := that.(type) {
	case akka.ActorPath:
		{
			return p.path.Equals(other)
		}
	}
	return false
}

func (p *ActorSystemImpl) Path() akka.ActorPath {
	return p.path
}

func (p *ActorSystemImpl) String() string {
	return ""
}

func (p *ActorSystemImpl) Tell(message interface{}, sender akka.ActorRef) {
	return
}

func (p *ActorSystemImpl) ActorOf(props akka.Props, name string) (ref akka.ActorRef, err error) {
	return p.guardian.Cell().AttachChild(props, name, false)
}

func (p *ActorSystemImpl) Stop(actor akka.ActorRef) (err error) {
	return
}

func (p *ActorSystemImpl) ActorSelection(path akka.ActorPath) (selection akka.ActorSelection, err error) {
	return
}

func (p *ActorSystemImpl) Provider() akka.ActorRefProvider {
	return nil
}

func (p *ActorSystemImpl) createDynamicAccess() dynamic_access.DynamicAccess {
	return dynamic_access.NewReflectiveDynamicAccess(p.classLoader)
}

func (p *ActorSystemImpl) configureSettings(config *configuration.Config) (err error) {
	var settings *akka.Settings
	if settings, err = akka.NewSettings(p.name, config); err != nil {
		return
	}

	p.settings = settings

	return err
}

func (p *ActorSystemImpl) configureEventStream() (err error) {
	// p.eventStream = NewEventStream(p.settings.DebugEventStream)
	return
}

func (p *ActorSystemImpl) configureLoggers() (err error) {
	return
}

func (p *ActorSystemImpl) configureScheduler() (err error) {
	schedulerType, exist := p.classLoader.ClassNameOf(p.settings.SchedulerClass)
	if !exist {
		err = akka.ErrTypeNotExistInClassLoader
		return
	}

	var ins interface{}
	ins, err = p.dynamicAccess.CreateInstanceByType(schedulerType, p.settings.Config(), nil)
	if err != nil {
		return
	}

	scheduler, ok := ins.(akka.Scheduler)
	if ok {
		p.scheduler = scheduler
		return
	}

	err = akka.ErrBadTypeOfScheduler

	return

}

func (p *ActorSystemImpl) configureProvider() (err error) {
	var obj interface{}
	if obj, err = p.dynamicAccess.CreateInstanceByName(p.settings.ProviderClass, p); err != nil {
		return
	}

	if provider, ok := obj.(akka.ActorRefProvider); !ok {
		err = akka.ErrCreateActorRefProviderFailure
		return
	} else {
		p.provider = provider
	}

	return
}

func (p *ActorSystemImpl) configureMailboxes() (err error) {
	return
}

func (p *ActorSystemImpl) configureDispatchers() {
	return
}
