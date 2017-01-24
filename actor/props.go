package actor

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/dispatch"
)

var (
	Props = &ActorProps{}
)

type ActorProps struct {
	deploy       akka.Deploy
	mailbox      string
	dispatcher   string
	routerConfig akka.RouterConfig
	producer     IndirectActorProducer
}

func (p ActorProps) Create(v interface{}, args ...interface{}) (props akka.Props, err error) {
	var producer IndirectActorProducer
	if producer, err = _CreateProducer(v, args...); err != nil {
		return
	}

	props = &ActorProps{
		producer:   producer,
		mailbox:    dispatch.DefaultMailboxId,
		dispatcher: dispatch.DefaultDispatcherId,
	}

	return
}

func (p ActorProps) NewActor() (actor akka.Actor, err error) {
	actor, err = p.producer.Produce()
	return
}

func (p ActorProps) Dispatcher() string {
	if len(p.deploy.Dispatcher()) == 0 {
		return dispatch.DefaultDispatcherId
	}
	return p.deploy.Dispatcher()
}

func (p ActorProps) Mailbox() string {
	if len(p.deploy.Mailbox()) == 0 {
		return dispatch.DefaultMailboxId
	}
	return p.deploy.Mailbox()
}

func (p ActorProps) RouterConfig() akka.RouterConfig {
	return p.deploy.RouterConfig()
}

func (p ActorProps) WithDeploy(deploy akka.Deploy) (props akka.Props, err error) {
	newProps := p.copy()
	newProps.deploy = deploy
	return
}

func (p ActorProps) WithDispatcher(dispatcher string) (props akka.Props, err error) {
	newProps := p.copy()
	newProps.dispatcher = dispatcher
	return
}

func (p ActorProps) WithMailbox(mailbox string) (props akka.Props, err error) {
	newProps := p.copy()
	newProps.mailbox = mailbox
	return
}

func (p ActorProps) WithRouter(config akka.RouterConfig) (props akka.Props, err error) {
	newProps := p.copy()
	newProps.routerConfig = config
	return
}

func (p ActorProps) copy() (props *ActorProps) {
	return &ActorProps{
		deploy:       p.deploy,
		mailbox:      p.mailbox,
		dispatcher:   p.dispatcher,
		routerConfig: p.routerConfig,
		producer:     p.producer,
	}
}
