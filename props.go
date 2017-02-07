package akka

import (
	"reflect"
)

type Props interface {
	Type() reflect.Type
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
