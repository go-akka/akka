package akka

import (
	"sort"
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
	uid    int64
}

func NewChildActorPath(parent ActorPath, name string, uid int64) ActorPath {

	return &ChildActorPath{
		parent: parent,
		name:   name,
		uid:    uid,
	}
}

func (p *ChildActorPath) Uid() int64 {
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
		sort.Sort(sort.Reverse(sort.StringSlice(elements)))
	}
	return elements
}

func (p *ChildActorPath) Name() (name string) {
	return p.name
}

func (p *ChildActorPath) Parent() (parent ActorPath) {
	return p.parent
}

func (p *ChildActorPath) Root() (root RootActorPath) {
	r := p.rootRec(p)
	return *r
}

func (p *ChildActorPath) rootRec(path ActorPath) *RootActorPath {
	switch v := path.(type) {
	case *RootActorPath:
		return v
	case *ChildActorPath:
		return v.rootRec(v.parent)
	default:
		return nil
	}
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

func (p *ChildActorPath) splitNameAndUid(name string) (n string, uid int64) {
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
	addr := p.Address()
	return addr.String() + "/" + strings.Join(p.Elements(), "/")
}
