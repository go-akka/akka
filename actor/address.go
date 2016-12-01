package actor

type Address struct {
}

func (p *Address) Host() string {
	return ""
}

func (p *Address) Port() int {
	return 0
}

func (p *Address) System() string {
	return ""
}

func (p *Address) Protocol() string {
	return ""
}
