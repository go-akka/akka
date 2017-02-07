package props

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/dispatch"
	"reflect"
)

var (
	emptyProps              = &Props{}
	globalProducerCreatorFn ProducerCreatorFunc
)

type ProducerCreatorFunc func(v interface{}, args ...interface{}) (producer IndirectActorProducer, err error)

func RegisterGlobalProducerCreator(fn ProducerCreatorFunc) {
	globalProducerCreatorFn = fn
	emptyProps.producerCreator = fn
}

func Create(v interface{}, args ...interface{}) (*Props, error) {

	var val interface{}
	if obj, ok := v.(reflect.Type); ok {
		for obj.Kind() == reflect.Ptr {
			obj = obj.Elem()
		}
		val = reflect.New(obj).Interface()
	} else {
		val = v
	}

	v, e := emptyProps.Create(val, args...)
	if e != nil {
		return nil, e
	}

	return v.(*Props), nil
}

type Props struct {
	deploy          akka.Deploy
	mailbox         string
	dispatcher      string
	routerConfig    akka.RouterConfig
	producer        IndirectActorProducer
	producerCreator ProducerCreatorFunc
	typ             reflect.Type
}

func (p Props) Create(v interface{}, args ...interface{}) (props akka.Props, err error) {
	var producer IndirectActorProducer
	if producer, err = createProducer(p.producerCreator, v, args...); err != nil {
		return
	}

	props = &Props{
		producer:        producer,
		producerCreator: p.producerCreator,
		mailbox:         dispatch.DefaultMailboxId,
		dispatcher:      dispatch.DefaultDispatcherId,
		typ:             reflect.TypeOf(v),
	}

	return
}

func (p Props) NewActor() (actor akka.Actor, err error) {
	actor, err = p.producer.Produce()
	return
}

func (p Props) Dispatcher() string {
	if len(p.deploy.Dispatcher()) == 0 {
		return dispatch.DefaultDispatcherId
	}
	return p.deploy.Dispatcher()
}

func (p Props) Mailbox() string {
	if len(p.deploy.Mailbox()) == 0 {
		return dispatch.DefaultMailboxId
	}
	return p.deploy.Mailbox()
}

func (p Props) RouterConfig() akka.RouterConfig {
	return p.deploy.RouterConfig()
}

func (p Props) WithDeploy(deploy akka.Deploy) (props akka.Props) {
	newProps := p.copy()
	newProps.deploy = deploy
	return newProps
}

func (p Props) WithDispatcher(dispatcher string) (props akka.Props) {
	newProps := p.copy()
	newProps.deploy = p.deploy.WithDispatcher(dispatcher)
	return newProps
}

func (p Props) WithMailbox(mailbox string) (props akka.Props) {
	newProps := p.copy()
	newProps.deploy = p.deploy.WithMailbox(mailbox)
	return newProps
}

func (p Props) WithRouter(config akka.RouterConfig) (props akka.Props) {
	newProps := p.copy()
	newProps.deploy = p.deploy.WithRouterConfig(config)
	return newProps
}

func (p Props) Type() reflect.Type {
	return p.typ
}

func (p Props) copy() (props *Props) {
	return &Props{
		deploy:          p.deploy,
		mailbox:         p.mailbox,
		dispatcher:      p.dispatcher,
		routerConfig:    p.routerConfig,
		producer:        p.producer,
		producerCreator: p.producerCreator,
		typ:             p.typ,
	}
}
