package actor

import (
	"fmt"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/orcaman/concurrent-map"

	"github.com/go-akka/akka"
	"github.com/go-akka/akka/dispatch"
	"github.com/go-akka/akka/event"
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
	mailboxes     akka.Mailboxes
	deadletters   akka.ActorRef
	dispatchers   akka.Dispatchers

	provider         akka.ActorRefProvider
	extensions       cmap.ConcurrentMap
	extensionsLocker sync.Mutex

	log akka.LoggingAdapter
}

func AkkaClassLoader() class_loader.ClassLoader {
	return class_loader.Default
}

func NewActorSystem(name string, config ...*configuration.Config) (system *ActorSystemImpl, err error) {
	if match, _ := regexp.MatchString("^[a-zA-Z0-9][a-zA-Z0-9-_]*$", name); !match {
		err = akka.ErrInvalidActorSystemName
		return
	}

	classLoader := class_loader.NewClassicClassLoader(class_loader.Default)

	sys := &ActorSystemImpl{
		name:          name,
		startedTime:   time.Now(),
		classLoader:   classLoader,
		dynamicAccess: dynamic_access.NewReflectiveDynamicAccess(classLoader),
		extensions:    cmap.New(),
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

	if err = sys.configureLoggers(); err != nil {
		return
	}

	// if err = sys.configureScheduler(); err != nil {
	// 	return
	// }
	sys.configureProvider()
	// sys.configureTerminationCallbacks()
	sys.configureMailboxes()
	sys.configureDispatchers()

	system = sys

	err = system.Start()

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

func (p *ActorSystemImpl) DynamicAccess() dynamic_access.DynamicAccess {
	return p.dynamicAccess
}

func (p *ActorSystemImpl) Provider() akka.ActorRefProvider {
	return p.provider
}

func (p *ActorSystemImpl) LogFilter() akka.LoggingFilter {
	val, _ := p.dynamicAccess.CreateInstanceByName(p.settings.LoggingFilter, p.settings, p.eventStream)
	return val.(akka.LoggingFilter)
}

func (p *ActorSystemImpl) Guardian() akka.LocalActorRef {
	return p.provider.Guardian()
}

func (p *ActorSystemImpl) SystemGuardian() akka.LocalActorRef {
	return p.provider.SystemGuardian()
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

func (p *ActorSystemImpl) Log() akka.LoggingAdapter {
	return p.log
}

func (p *ActorSystemImpl) DeadLetters() akka.ActorRef {
	return p.deadletters
}

func (p *ActorSystemImpl) EventStream() akka.EventStream {
	return p.eventStream
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
			return p.path.CompareTo(other) == 0
		}
	}
	return false
}

func (p *ActorSystemImpl) Path() akka.ActorPath {
	return p.path
}

func (p *ActorSystemImpl) String() string {
	return p.LookupRoot().Path().Root().Address().String()
}

func (p *ActorSystemImpl) Tell(message interface{}, sender akka.ActorRef) {
	return
}

func (p *ActorSystemImpl) ActorOf(props akka.Props, name string) (ref akka.ActorRef, err error) {
	return p.Guardian().Underlying().AttachChild(props, name, false)
}

func (p *ActorSystemImpl) SystemActorOf(props akka.Props, name string) (ref akka.ActorRef, err error) {
	return p.SystemGuardian().Underlying().AttachChild(props, name, true)
}

func (p *ActorSystemImpl) Stop(actor akka.ActorRef) (err error) {
	// path := actor.Path()
	// guard := p.Guardian().Path()
	// sys := p.SystemGuardian().Path()

	return nil
}

func (p *ActorSystemImpl) ActorSelection(path akka.ActorPath) (selection akka.ActorSelection, err error) {
	return
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
	p.eventStream = event.NewEventStream(p, p.settings.DebugEventStream)
	p.eventStream.StartStdoutLogger(p.settings)
	return
}

func (p *ActorSystemImpl) configureLoggers() (err error) {
	p.log = event.NewBusLogging(p.eventStream, "ActorSystem("+p.name+")", reflect.TypeOf(p), &event.DefaultLogMessageFormatter{})

	return nil
}

func (p *ActorSystemImpl) configureScheduler() (err error) {
	schedulerType, exist := p.classLoader.ClassNameOf(p.settings.SchedulerClass)
	if !exist {
		err = fmt.Errorf("type not in class loader, %s: %s", "akka.scheduler.implementation", p.settings.SchedulerClass)
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
	if obj, err = p.dynamicAccess.CreateInstanceByName(p.settings.ProviderClass, p.name, p.settings, p.eventStream, p.dynamicAccess); err != nil {
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
	p.mailboxes = dispatch.NewMailboxes(p.settings, p.eventStream, p.dynamicAccess, p.deadletters)
	return
}

func (p *ActorSystemImpl) configureDispatchers() {
	p.dispatchers = dispatch.NewDispatchers(p.settings, dispatch.NewDefaultDispatcherPrerequisites(p.eventStream, p.scheduler, p.dynamicAccess, p.settings, p.mailboxes))
}

func (p *ActorSystemImpl) Start() (err error) {
	if err = p.provider.Init(p); err != nil {
		return
	}

	p.loadExtensions()

	return
}

func (p *ActorSystemImpl) LookupRoot() akka.InternalActorRef {
	return p.provider.RootGuardian()
}

func (p *ActorSystemImpl) loadExtensions() {

	loadExtensions := func(key string, throwOnLoadFail bool) {
		for _, extensionFqn := range p.settings.Config().GetStringList(key) {
			extensionType, exist := p.classLoader.ClassNameOf(extensionFqn)
			if !exist {
				p.log.Error(nil, "Extension of %s is not exist", extensionFqn)
				if throwOnLoadFail {
					panic(fmt.Errorf("Extension of %s is not exist", extensionFqn))
				}
				continue
			}

			ins, err := p.dynamicAccess.CreateInstanceByType(extensionType)
			if err != nil {
				p.log.Error(err, "While trying to load extension %s skipping...", extensionFqn)
				if throwOnLoadFail {
					panic(fmt.Errorf("While trying to load extension %s, %s", extensionFqn, err))
				}
				continue
			}

			switch v := ins.(type) {
			case akka.ExtensionProvider:
				{
					p.RegisterExtension(v.Lookup())
				}
			case akka.ExtensionId:
				{
					p.RegisterExtension(v)
				}
			}
		}
	}

	loadExtensions("akka.library-extensions", true)
	loadExtensions("akka.extensions", false)
}

func (p *ActorSystemImpl) RegisterExtension(ext akka.ExtensionId) akka.Extension {
	if ext == nil {
		return nil
	}

	p.extensionsLocker.Lock()
	defer p.extensionsLocker.Unlock()

	extensionIns, exist := p.extensions.Get(ext.ExtensionType().String())
	if exist {
		return extensionIns.(akka.Extension)
	}

	newextensionIns := ext.CreateExtension(p)
	p.extensions.Set(ext.ExtensionType().String(), newextensionIns)

	return newextensionIns
}

func (p *ActorSystemImpl) Extension(ext akka.ExtensionId) akka.Extension {
	v, exist := p.extensions.Get(ext.ExtensionType().String())
	if exist {
		return v.(akka.Extension)
	}
	return nil
}

func (p *ActorSystemImpl) HasExtension(ext akka.ExtensionId) bool {
	_, exist := p.extensions.Get(ext.ExtensionType().String())
	return exist
}
