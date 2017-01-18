package akka

type Mailboxes interface {
	Lookup(id string) (t MailboxType, exist bool)
}
