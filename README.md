Go AKKA
=======

### Currently Not Available ！！！

`example.go`

```go
func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()

	config := configuration.LoadConfig("akka.conf")

	var system ActorSystem
	if system, err = actor.NewActorSystem("test", config); err != nil {
		return
	}

	var props Props
	if props, err = actor.Props.Create((*HelloWorld)(nil)); err != nil {
		return
	}

	var helloWorld ActorRef
	if helloWorld, err = system.ActorOf(props, "HelloWorld"); err != nil {
		return
	}

	var terminatorProps Props
	if terminatorProps, err = actor.Props.Create((*Terminator)(nil), helloWorld); err != nil {
		return
	}

	if _, err = system.ActorOf(terminatorProps, "Terminator"); err != nil {
		return
	}

	fmt.Println("Press return to exit ...")
	fmt.Scanln()
}
```

`akka.conf`

```hocon
akka {
  log-config-on-start = on
  stdout-loglevel = DEBUG
  loglevel = ERROR
  actor {
      provider = "LocalActorRefProvider"

      default-mailbox {
        mailbox-type = "akka.dispatch.unbounded-mailbox"
      }

      default-dispatcher {
        type = "dispatcher"
      }
 
      debug {  
        receive = on 
        autoreceive = on
        lifecycle = on
        event-stream = on
        unhandled = on
      }
  }
}
```


`HelloWorld`

```go
type HelloWorld struct {
	*actor.UntypedActor
}

func (p *HelloWorld) PreStart() (err error) {
	var props Props
	props, err = actor.Props.Create((*Greeter)(nil))
	if err != nil {
		fmt.Println(err)
		return
	}

	var greeter ActorRef
	greeter, err = p.Context().ActorOf(props, "greeter")
	if err != nil {
		fmt.Println(err)
		return
	}

	greeter.Tell("hello greeter", p.Context().Self())

	return
}

func (p *HelloWorld) Receive(message interface{}) (unhandled bool, err error) {
	switch msg := message.(type) {
	case string:
		{
			p.Context().StopActor(p.Self())
		}
	default:
		unhandled = true
	}

	return
}
```


`Greeter`

```go
type Greeter struct {
	*actor.UntypedActor
}

func (p *Greeter) PreStart() {
	fmt.Println("pre start at Greeter")
}

func (p *Greeter) Receive(message interface{}) (unhandled bool, err error) {

	switch msg := message.(type) {
	case string:
		{
			fmt.Println("Greeter received message:", msg)
		}
	default:
		unhandled = true
	}

	return
}
```


`Terminator`

```go
type Terminator struct {
	*actor.UntypedActor

	ref ActorRef
}

func (p *Terminator) Terminator(ref ActorRef) (err error) {
	p.ref = ref
	p.Context().Watch(ref)
	return
}

func (p *Terminator) Receive(message interface{}) (unhandled bool, err error) {
	switch message.(type) {
	case Terminated:
		{
			fmt.Printf("%s has terminated, shutting down system", p.ref.Path())
			p.Context().System().Terminate()
		}
	default:
		unhandled = true
	}
	return
}
```