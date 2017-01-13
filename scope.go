package akka

type Scope struct {
}

func (p *Scope) WithFallback(other Scope) (scope Scope) {
	return
}
