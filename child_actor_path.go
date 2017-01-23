package akka

import (
	"strconv"
	"strings"
)

var (
	_ ActorPath = (*ChildActorPath)(nil)
)

func ActorPathFromString(path string) (actorPath ActorPath, err error) {
	return
}

type ChildActorPath struct {
	parent ActorPath
	name   string
	uid    int
}

func NewChildActorPath(parent ActorPath, name string, uid int) ActorPath {

	return &ChildActorPath{
		parent: parent,
		name:   name,
		uid:    uid,
	}
}

func (p *ChildActorPath) Uid() int {
	return p.uid
}

func (p *ChildActorPath) Address() (addr *Address) {
	return
}

func (p *ChildActorPath) Elements() (elems []string) {
	return
}

func (p *ChildActorPath) Name() (name string) {
	return
}

func (p *ChildActorPath) Parent() (parent ActorPath) {
	return
}

func (p *ChildActorPath) Root() (root *RootActorPath) {
	return
}

func (p *ChildActorPath) CompareTo(other ActorPath) int {
	return 0
}

func (p *ChildActorPath) ToSerializationFormat() string {
	return ""
}

func (p *ChildActorPath) ToSerializationFormatWithAddress(address Address) string {
	return ""
}

func (p *ChildActorPath) ToStringWithAddress(address Address) string {
	return ""
}

func (p *ChildActorPath) splitNameAndUid(name string) (n string, uid int) {
	i := strings.Index(name, "#")
	if i < 0 {
		n = name
		return
	}

	n = name[0:i]
	uid, _ = strconv.Atoi(name[i+1:])
	return
}

func (p *ChildActorPath) Append(name string) ActorPath {
	childName, uid := p.splitNameAndUid(name)
	return NewChildActorPath(p, childName, uid)
}

func (p *ChildActorPath) ToStringWithoutAddress() string {
	return ""
}

func (p *ChildActorPath) Child(child string) (path ActorPath, err error) {
	return
}

func (p *ChildActorPath) Descendant(names []string) (path ActorPath, err error) {
	return
}

func (p *ChildActorPath) String() string {
	return ""
}
