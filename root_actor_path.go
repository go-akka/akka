package akka

import (
	"strconv"
	"strings"
)

var (
	_ ActorPath = (*RootActorPath)(nil)
)

type RootActorPath struct {
	name    string
	address Address
}

func NewRootActorPath(address Address, name string) ActorPath {

	if len(name) == 0 {
		name = "/"
	}

	return &RootActorPath{
		name:    name,
		address: address,
	}
}

func (p *RootActorPath) Uid() int64 {
	return 0
}

func (p *RootActorPath) Address() (addr Address) {
	return p.address
}

func (p *RootActorPath) Elements() (elems []string) {
	return
}

func (p *RootActorPath) Name() (name string) {
	return p.name
}

func (p *RootActorPath) Parent() (parent ActorPath) {
	return
}

func (p *RootActorPath) Root() (root RootActorPath) {
	return
}

func (p *RootActorPath) CompareTo(other ActorPath) int {
	return 0
}

func (p *RootActorPath) ToSerializationFormat() string {
	return ""
}

func (p *RootActorPath) ToSerializationFormatWithAddress(address Address) string {
	return ""
}

func (p *RootActorPath) ToStringWithAddress(address Address) string {
	return ""
}

func (p *RootActorPath) ToStringWithoutAddress() string {
	return ""
}

func (p *RootActorPath) Child(child string) (path ActorPath, err error) {
	return
}

func (p *RootActorPath) Descendant(names []string) (path ActorPath, err error) {
	return
}

func (p *RootActorPath) splitNameAndUid(name string) (n string, uid int64) {
	i := strings.Index(name, "#")
	if i < 0 {
		n = name
		return
	}

	n = name[0:i]
	v, _ := strconv.Atoi(name[i+1:])
	uid = int64(v)
	return
}

func (p *RootActorPath) Append(name string) ActorPath {
	childName, uid := p.splitNameAndUid(name)
	return NewChildActorPath(p, childName, uid)
}

func (p *RootActorPath) String() string {
	return p.address.String() + p.name
}
