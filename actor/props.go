package actor

import (
	"github.com/go-akka/akka"
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
		producer: producer,
	}

	return
}

func (p ActorProps) newActor() (actor akka.Actor, err error) {
	actor, err = p.producer.Produce()
	return
}

func (p ActorProps) Dispatcher() string {
	return p.deploy.Dispatcher()
}

func (p ActorProps) Mailbox() string {
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
