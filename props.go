package akka

type Props interface {
	NewActor() (actor Actor, err error)
	Create(v interface{}, args ...interface{}) (props Props, err error)
	Dispatcher() string
	Mailbox() string
	RouterConfig() RouterConfig
	WithDeploy(deploy Deploy) (props Props)
	WithDispatcher(dispatcher string) (props Props)
	WithMailbox(mailbox string) (props Props)
	WithRouter(config RouterConfig) (props Props)
}
