package akka

type MessageInvoker interface {
	SystemInvoke(message SystemMessage)
	Invoke(envelop Envelope)
}
