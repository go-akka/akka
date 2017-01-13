package actor

import (
	"errors"
	"github.com/go-akka/akka"
	"testing"
)

type testUntypedActor struct {
	*UntypedActor

	inited bool
}

func (p *testUntypedActor) Init(arg1, arg2 int) error {
	p.inited = true

	return nil
}

func (p *testUntypedActor) Receive(message interface{}) (handled bool, err error) {
	if p.inited == false {
		err = errors.New("testUntypedActor args init failure")
		return
	}
	handled = true
	return
}

type testReceiveActor struct {
	*ReceiveActor

	inited         bool
	stringReceived bool
	intReceived    bool
}

func (p *testReceiveActor) Init(arg1, arg2 int) error {
	p.inited = true

	p.SmartReceive(func(message string) {
		p.stringReceived = true
	})

	p.SmartReceive(func(message int) {
		p.intReceived = true
	})

	return nil
}

func TestCreateUntypedActor(t *testing.T) {
	var err error
	var producer IndirectActorProducer

	producer, err = _CreateProducer((*testUntypedActor)(nil), 1, 2)

	if err != nil {
		t.Fatalf("producer create failure: %s", err.Error())
		return
	}

	var actor akka.Actor
	if actor, err = producer.Produce(); err != nil {
		t.Fatalf("produce actor failure: %s", err.Error())
		return
	}

	if handled, _ := actor.Receive("hello"); !handled {
		t.Fatalf("testUntypedActor Receive unhandled")
		return
	}
}

func TestCreateReciveActor(t *testing.T) {
	var err error
	var producer IndirectActorProducer

	producer, err = _CreateProducer((*testReceiveActor)(nil), 1, 2)

	if err != nil {
		t.Fatalf("producer create failure: %s", err.Error())
		return
	}

	var actor akka.Actor
	if actor, err = producer.Produce(); err != nil {
		t.Fatalf("produce actor failure: %s", err.Error())
		return
	}

	if handled, _ := actor.Receive("hello"); !handled {
		t.Fatalf("testReceiveActor Receive unhandled")
		return
	}

	if handled, _ := actor.Receive(32); !handled {
		return
	}

}
