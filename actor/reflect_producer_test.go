package actor

import (
	"errors"
	"testing"
)

type testActor struct {
	*UntypedActor

	inited bool
}

func (p *testActor) Init(arg1, arg2 int) error {
	p.inited = true

	return nil
}

func (p *testActor) Receive(message interface{}) (handled bool, err error) {
	if p.inited == false {
		err = errors.New("testActor args init failure")
		return
	}
	handled = true
	return
}

func TestCreateReflectProducer(t *testing.T) {
	var err error
	var producer IndirectActorProducer

	producer, err = _CreateProducer((*testActor)(nil), 1, 2)

	if err != nil {
		t.Fatalf("producer create failure: %s", err.Error())
		return
	}

	var actor Actor
	if actor, err = producer.Produce(); err != nil {
		t.Fatalf("produce actor failure: %s", err.Error())
		return
	}

	if handled, _ := actor.Receive("hello"); !handled {
		t.Fatalf("testActor Receive unhandled")
		return
	}

}
