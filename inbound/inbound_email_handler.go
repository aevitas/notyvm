package inbound

import (
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
	"github.com/patrickmn/go-cache"
	"github.com/sendgrid/sendgrid-go/helpers/inbound"
)

type email struct {
	Id         uint64   `json:"id"`
	Sender     string   `json:"sender"`
	Recipients []string `json:"recipients"`
	SenderName string   `json:"sender_name"`
	Subject    string   `json:"subject"`
	Text       string   `json:"text"`
}

func ProcessInboundEmail(ctx *gin.Context, cache *cache.Cache) error {
	parsed, err := inbound.Parse(ctx.Request)

	if err != nil {
		log.Fatal(err)
	}

	msg := email{
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

	if len(msg.Recipients) > 0 {
		return errors.New("ignoring message with multiple recipients; we don't listen to cc or bcc")
	}

	for _, m := range msg.Recipients {
		if !strings.Contains("@isveiled.com", m) && !strings.Contains("@veiled.io", m) {
			log.Printf("%s is not addressed to us; ignoring", m)
			continue
		}

		inbox := make(map[string]map[uint64]email)
		ib, f := cache.Get(m)

		if f {
			inbox = ib.(map[string]map[uint64]email)
			log.Printf("inbox for %s retrieved from cache", m)
		}

		inbox[m][msg.Id] = msg

		log.Printf("delivered message %d to %s", msg.Id, m)
	}

	return nil
}
