package actor

type DeployOption func(d *Deploy)

type Deploy struct {
	scope        Scope
	routerConfig RouterConfig
	path         string
	mailbox      string
	dispatcher   string
	config       Config
}

func (p Deploy) WithFallback(other Deploy) (deploy Deploy, err error) {
	return
}

func NewDeploy(opts ...DeployOption) Deploy {
	deploy := Deploy{
		path: "",
	}

	for _, opt := range opts {
		opt(&deploy)
	}

	return deploy
}

func DeployConfig(config Config) DeployOption {
	return func(d *Deploy) {
		d.config = config
	}
}

func DeployRouterConfig(config RouterConfig) DeployOption {
	return func(d *Deploy) {
		d.routerConfig = config
	}
}

func DeployScope(scope Scope) DeployOption {
	return func(d *Deploy) {
		d.scope = scope
	}
}

func DeployDispatcher(dispatcher string) DeployOption {
	return func(d *Deploy) {
		d.dispatcher = dispatcher
	}
}

func DeployMailbox(mailbox string) DeployOption {
	return func(d *Deploy) {
		d.mailbox = mailbox
	}
}
