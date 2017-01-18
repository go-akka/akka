package actor

import (
	"github.com/go-akka/akka"
)

type BehaviorStack struct {
	values []akka.ReceiveFunc
}

func NewBehaviorStack() *BehaviorStack {
	return &BehaviorStack{make([]akka.ReceiveFunc, 0)}
}

func (p *BehaviorStack) Push(v akka.ReceiveFunc) {
	p.values = append(p.values, v)
}

func (p *BehaviorStack) Pop() (recv akka.ReceiveFunc, exist bool) {
	l := len(p.values)
	if l == 0 {
		return
	}

	res := p.values[l-1]
	p.values[l-1] = nil
	p.values = p.values[:l-1]

	recv = res
	exist = true
	return
}

func (p *BehaviorStack) Current() (recv akka.ReceiveFunc, exist bool) {
	l := len(p.values)
	if l == 0 {
		return
	}

	recv = p.values[l-1]
	exist = true
	return
}

func (p *BehaviorStack) Len() int {
	return len(p.values)
}
