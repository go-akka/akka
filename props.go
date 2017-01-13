package akka

type Props interface {
	Create(v interface{}, args ...interface{}) (props Props, err error)
	Dispatcher() string
	Mailbox() string
	RouterConfig() RouterConfig
	WithDeploy(deploy Deploy) (props Props, err error)
	WithDispatcher(dispatcher string) (props Props, err error)
	WithMailbox(mailbox string) (props Props, err error)
	WithRouter(config RouterConfig) (props Props, err error)
}
