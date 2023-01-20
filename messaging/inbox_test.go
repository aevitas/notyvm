package messaging

import (
	"testing"

	"aevitas.dev/veiled/models"
)

func TestInboxMessages(t *testing.T) {
	ib := Inbox{
		Messages: map[uint64]models.Email{},
	}

	ib.AddMessage(models.Email{Id: 10, Sender: "test@veiled.io", Subject: "hello"})

	msgs := ib.ListMessages()

	for _, m := range msgs {
		if m.Subject != "hello" {
			t.Fail()
		}
		if m.Id != 10 {
			t.Fail()
		}
	}
}
