package akka

import (
	"bytes"
	"strconv"
)

type Address struct {
	host     string
	port     int
	system   string
	protocol string
}

func NewAddress(protocol string, system string, host string, port int) Address {
	return Address{
		protocol: protocol,
		system:   system,
		host:     host,
		port:     port,
	}
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

func (p *Address) HasGlobalScope() bool {
	return len(p.host) > 0
}

func (p *Address) HasLocalScope() bool {
	return len(p.host) == 0
}

func (p Address) String() string {
	if p.protocol != "akka" {
		panic("")
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString(p.protocol)
	buf.WriteString("://")
	buf.WriteString(p.system)

	if len(p.host) > 0 {
		buf.WriteByte('@')
		buf.WriteString(p.host)
	}

	if p.port > 0 {
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(p.port))
	}

	return buf.String()
}

func (p *Address) HostPort() string {
	buf := bytes.NewBuffer(nil)
	if len(p.host) > 0 {
		buf.WriteByte('@')
		buf.WriteString(p.host)
	}

	if p.port > 0 {
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(p.port))
	}

	return buf.String()
}
