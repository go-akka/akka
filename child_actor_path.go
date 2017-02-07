package akka

import (
	"fmt"
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

func (p *ChildActorPath) Address() (addr Address) {
	addr = p.Root().address
	return
}

func (p *ChildActorPath) Elements() (elems []string) {
	var current ActorPath = p
	var elements = []string{}
	for {
		if _, ok := current.(*RootActorPath); ok {
			break
		}

		elements = append(elements, current.Name())
		current = current.Parent()
	}
	elements = p.reverse(elements)
	return elements
}

func (p *ChildActorPath) Name() (name string) {
	return p.name
}

func (p *ChildActorPath) Parent() (parent ActorPath) {
	return p.parent
}

func (p *ChildActorPath) Root() (root *RootActorPath) {
	current := p.parent
	for {
		if currentC, ok := current.(*ChildActorPath); ok {
			current = currentC.parent
			continue
		}
		break
	}

	return current.Root()
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
	return address.String() + p.Join()
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
	return p.Join()
}

func (p *ChildActorPath) Child(child string) (path ActorPath, err error) {
	return
}

func (p *ChildActorPath) Descendant(names []string) (path ActorPath, err error) {
	return
}

func (p *ChildActorPath) Join() string {
	joined := strings.Join(p.Elements(), "/")
	return "/" + joined
}

func (p *ChildActorPath) String() string {

	if p.Uid() == 0 {
		return p.ToStringWithAddress(p.Address())
	}
	return fmt.Sprintf("%s#%d", p.ToStringWithAddress(p.Address()), p.Uid())
}

func (p *ChildActorPath) reverse(values []string) []string {
	lv := len(values)
	newV := make([]string, lv)
	for i := 0; i < lv; i++ {
		newV[lv-1-i] = values[i]
	}

	return newV
}
