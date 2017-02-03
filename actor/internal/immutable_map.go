package internal

import (
	"github.com/orcaman/concurrent-map"
)

type ImmutableMap struct {
	cmap.ConcurrentMap
}

func NewImmutableMap() *ImmutableMap {
	return &ImmutableMap{
		cmap.New(),
	}
}

func (p *ImmutableMap) Add(k string, v interface{}) *ImmutableMap {
	if p.Has(k) {
		return p
	}

	cp := p.copy()
	if cp.ConcurrentMap.SetIfAbsent(k, v) {
		return cp
	}
	return p
}

func (p *ImmutableMap) Set(k string, v interface{}) *ImmutableMap {
	cp := p.copy()
	cp.ConcurrentMap.Set(k, v)
	return cp
}

func (p *ImmutableMap) Remove(k string) *ImmutableMap {
	cp := p.copy()
	cp.ConcurrentMap.Remove(k)
	return cp
}

func (p *ImmutableMap) Get(k string) (v interface{}, exist bool) {
	return p.ConcurrentMap.Get(k)
}

func (p *ImmutableMap) copy() *ImmutableMap {
	newMap := cmap.New()

	for k, v := range p.ConcurrentMap.Items() {
		newMap.Set(k, v)

	}

	return &ImmutableMap{
		newMap,
	}
}
