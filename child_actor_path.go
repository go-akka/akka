package akka

var (
	_ ActorPath = (*ChildActorPath)(nil)
)

func ActorPathFromString(path string) (actorPath ActorPath, err error) {
	return
}

type ChildActorPath struct {
	uid int
}

func NewChildActorPath(parent ActorPath, name string) (path ActorPath, err error) {
	return
}

func (p *ChildActorPath) Uid() int {
	return p.uid
}

func (p *ChildActorPath) Address() (addr Address) {
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

func (p *ChildActorPath) Root() (root RootActorPath) {
	return
}

func (p *ChildActorPath) Equals(other ActorPath) bool {
	return false
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

func (p *ChildActorPath) ToStringWithoutAddress() string {
	return ""
}

func (p *ChildActorPath) Child(child string) (path ActorPath, err error) {
	return
}

func (p *ChildActorPath) Descendant(names []string) (path ActorPath, err error) {
	return
}
