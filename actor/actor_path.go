package actor

type ActorPath struct {
	path string
	UID  string
}

func (p ActorPath) Equals(path ActorPath) bool {
	if path.UID == p.UID {
		return true
	}

	return false
}

func (p *ActorPath) String() string {
	return p.path
}

func ActorPathFromString(path string) (actorPath ActorPath, err error) {
	return
}
