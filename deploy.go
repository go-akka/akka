package akka

import (
	"github.com/go-akka/configuration"
)

type Deploy struct {
	scope        Scope
	routerConfig RouterConfig
	path         string
	mailbox      string
	dispatcher   string
	config       *configuration.Config
}

func (p Deploy) WithFallback(other Deploy) (deploy Deploy) {

	newDeploy := Deploy{
		dispatcher: p.dispatcher,
		mailbox:    p.mailbox,
	}

	if p.config != other.config {
		newDeploy.config = p.config.WithFallback(other.config)
	}

	newDeploy.routerConfig = p.routerConfig.WithFallback(other.routerConfig)
	newDeploy.scope = p.scope.WithFallback(other.scope)

	if len(p.dispatcher) == 0 {
		newDeploy.dispatcher = other.dispatcher
	}

	if len(p.mailbox) == 0 {
		newDeploy.mailbox = other.mailbox
	}

	deploy = newDeploy

	return
}

func (p Deploy) WithScope(scope Scope) Deploy {
	deploy := p.copy()
	deploy.scope = scope
	return deploy
}

func (p Deploy) WithMailbox(mailbox string) Deploy {
	deploy := p.copy()
	deploy.mailbox = mailbox
	return deploy
}

func (p Deploy) WithDispatcher(dispatcher string) Deploy {
	deploy := p.copy()
	deploy.dispatcher = dispatcher
	return deploy
}

func (p *Deploy) WithRouterConfig(routerConfig RouterConfig) Deploy {
	deploy := p.copy()
	deploy.routerConfig = routerConfig
	return deploy
}

func (p Deploy) Dispatcher() string {
	return p.dispatcher
}

func (p Deploy) Mailbox() string {
	return p.mailbox
}

func (p Deploy) RouterConfig() RouterConfig {
	return p.routerConfig
}

func (p Deploy) copy() Deploy {
	return Deploy{
		scope:        p.scope,
		routerConfig: p.routerConfig,
		path:         p.path,
		mailbox:      p.mailbox,
		dispatcher:   p.dispatcher,
		config:       p.config,
	}
}
