package akka

var (
	_ ActorPath = (*RootActorPath)(nil)
)

type RootActorPath struct {
	uid int
}

func NewRootActorPath(address Address, name string) (path ActorPath, err error) {
	return
}

func (p *RootActorPath) Uid() int {
	return p.uid
}

func (p *RootActorPath) Address() (addr Address) {
	return
}

func (p *RootActorPath) Elements() (elems []string) {
	return
}

func (p *RootActorPath) Name() (name string) {
	return
}

func (p *RootActorPath) Parent() (parent ActorPath) {
	return
}

func (p *RootActorPath) Root() (root RootActorPath) {
	return
}

func (p *RootActorPath) Equals(other ActorPath) bool {
	return false
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
