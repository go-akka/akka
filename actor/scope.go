package actor

type Scope struct {
}

func (p *Scope) WithFallback(other Scope) (scope Scope, err error) {
	return
}
