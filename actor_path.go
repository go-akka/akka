package akka

type ActorPath interface {
	Uid() int64
	Address() (addr Address)
	Elements() (elems []string)
	Name() (name string)
	Parent() (parent ActorPath)
	Root() (root *RootActorPath)
	CompareTo(other ActorPath) int
	ToSerializationFormat() string
	ToSerializationFormatWithAddress(address Address) string
	ToStringWithAddress(address Address) string
	ToStringWithoutAddress() string
	Child(child string) (path ActorPath, err error)
	Descendant(names []string) (path ActorPath, err error)
	Append(name string) ActorPath
	String() string
}
