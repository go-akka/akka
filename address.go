package akka

type Address struct {
	host     string
	port     int
	system   string
	protocol string
}

func (p *Address) Host() string {
	return p.host
}

func (p *Address) Port() int {
	return p.port
}

func (p *Address) System() string {
	return p.system
}

func (p *Address) Protocol() string {
	return p.protocol
}
