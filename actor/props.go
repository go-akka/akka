package actor

var (
	_ Actor = (*emptyActor)(nil)
)

var (
	ActorProps = EmptyProps()
)

type emptyActor struct {
}

func (p *emptyActor) Receive(message interface{}) (unmatched bool) {
	return
}

type Props struct {
	actor  Actor
	deploy Deploy
}

func EmptyProps() Props {
	return Props{actor: &emptyActor{}}
}

func (p Props) Create(v interface{}, args ...interface{}) (props Props, err error) {
	return
}

func (p Props) Dispatcher() string {
	return p.deploy.dispatcher
}

func (p Props) Mailbox() string {
	return p.deploy.mailbox
}

func (p Props) RouterConfig() RouterConfig {
	return p.deploy.routerConfig
}

func (p Props) WithDeploy(deploy Deploy) (props Props, err error) {
	return
}

func (p Props) WithDispatcher(dispatcher string) (props Props, err error) {
	return
}

func (p Props) WithMailbox(mailbox string) (props Props, err error) {
	return
}

func (p Props) WithRouter(config RouterConfig) (props Props, err error) {
	return
}
