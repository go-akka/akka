package dispatch

import (
	"github.com/go-akka/akka"
	"github.com/go-akka/akka/pkg/dynamic_access"
	"github.com/go-akka/configuration"
	"github.com/orcaman/concurrent-map"
)

const (
	DefaultMailboxId = "akka.actor.default-mailbox"
)

type Mailboxes struct {
	settings      *akka.Settings
	eventStream   akka.EventStream
	deadLetters   akka.ActorRef
	dynamicAccess dynamic_access.DynamicAccess

	mailboxTypeConfigurators cmap.ConcurrentMap

	defaultMailboxConfig *configuration.Config
	deadLetterMailbox    akka.Mailbox
}

func NewMailboxes(
	settings *akka.Settings,
	eventStream akka.EventStream,
	dynamicAccess dynamic_access.DynamicAccess,
	deadLetters akka.ActorRef,
) akka.Mailboxes {

	mailboxes := &Mailboxes{
		settings:                 settings,
		eventStream:              eventStream,
		deadLetters:              deadLetters,
		dynamicAccess:            dynamicAccess,
		mailboxTypeConfigurators: cmap.New(),
		defaultMailboxConfig:     settings.Config().GetConfig(DefaultMailboxId),
	}

	return mailboxes
}

func (p *Mailboxes) Lookup(id string) (t akka.MailboxType, exist bool) {
	return p.lookupConfigurator(id)
}

func (p *Mailboxes) lookupConfigurator(id string) (t akka.MailboxType, exist bool) {
	v, ok := p.mailboxTypeConfigurators.Get(id)
	if !ok {
		if id == "unbounded" {
			t = NewUnboundedMailbox()
			exist = true
		} else {
			mailboxTypeName := p.config(id).GetString("mailbox-type")
			if ins, err := p.dynamicAccess.CreateInstanceByName(mailboxTypeName, p.settings, p.config(id)); err != nil {
				return
			} else {
				t = ins.(akka.MailboxType)
				exist = true
			}
		}

		p.mailboxTypeConfigurators.SetIfAbsent(id, t)
		return
	}

	t = v.(akka.MailboxType)

	return
}

func (p *Mailboxes) config(id string) *configuration.Config {
	return configuration.ParseString("id:" + id).WithFallback(p.settings.Config().GetConfig(id)).WithFallback(p.defaultMailboxConfig)
}
