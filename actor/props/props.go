package props

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/dispatch"
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
	v, e := emptyProps.Create(v, args...)
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
	return
}

func (p Props) WithDispatcher(dispatcher string) (props akka.Props) {
	newProps := p.copy()
	newProps.dispatcher = dispatcher
	return
}

func (p Props) WithMailbox(mailbox string) (props akka.Props) {
	newProps := p.copy()
	newProps.mailbox = mailbox
	return
}

func (p Props) WithRouter(config akka.RouterConfig) (props akka.Props) {
	newProps := p.copy()
	newProps.routerConfig = config
	return
}

func (p Props) copy() (props *Props) {
	return &Props{
		deploy:          p.deploy,
		mailbox:         p.mailbox,
		dispatcher:      p.dispatcher,
		routerConfig:    p.routerConfig,
		producer:        p.producer,
		producerCreator: p.producerCreator,
	}
}
