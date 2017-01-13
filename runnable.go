package akka

type Runnable interface {
	Run()
}

type Action interface {
	Action()
}
