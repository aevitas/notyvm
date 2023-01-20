package inbound

import (
	"errors"
	"log"
	"strings"
	"time"

	"aevitas.dev/veiled/messaging"
	"aevitas.dev/veiled/models"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
	"github.com/patrickmn/go-cache"
	"github.com/sendgrid/sendgrid-go/helpers/inbound"
)

func ProcessInboundEmail(ctx *gin.Context, cache *cache.Cache) error {
	parsed, err := inbound.Parse(ctx.Request)

	if err != nil {
		log.Fatal(err)
	}

	msg := models.Email{
		Id:         ulid.Now(),
		Sender:     parsed.Envelope.From,
		Recipients: parsed.Envelope.To,
		SenderName: parsed.Headers["From"],
		Text:       parsed.TextBody,
		Subject:    parsed.ParsedValues["subject"],
	}

	dkim := strings.Contains(parsed.ParsedValues["dkim"], "pass")

	if !dkim {
		return errors.New("spam check didn't pass")
	}

	if len(msg.Recipients) > 1 {
		return errors.New("ignoring message with multiple recipients; we don't listen to cc or bcc")
	}

	for _, m := range msg.Recipients {
		if !strings.Contains(m, "@isveiled.com") && !strings.Contains(m, "@veiled.io") {
			log.Printf("%s is not addressed to us; ignoring", m)
			continue
		}

		inbox := messaging.EmptyInbox()
		ib, f := cache.Get(m)

		if f {
			inbox = ib.(messaging.Inbox)
			log.Printf("inbox for %s retrieved from cache", m)
		}

		inbox.Messages[msg.Id] = msg

		cache.Set(m, inbox, 60*time.Minute)

		log.Printf("delivered message %d to %s", msg.Id, m)
	}

	return nil
}
