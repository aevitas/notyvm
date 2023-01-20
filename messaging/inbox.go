package messaging

import (
	"fmt"
	"sort"

	"aevitas.dev/veiled/models"
)

type Inbox struct {
	Messages map[uint64]models.Email
}

func (inbox *Inbox) ListMessages() []models.Email {
	r := make([]models.Email, len(inbox.Messages))

	for _, m := range inbox.Messages {
		r = append(r, m)
	}

	sort.Slice(r[:], func(i, j int) bool {
		return r[i].Id > r[j].Id
	})

	return r
}

func (inbox *Inbox) GetMessage(id uint64) *models.Email {
	m := inbox.Messages[id]

	return &m
}

func (inbox *Inbox) AddMessage(msg models.Email) error {

	if inbox.GetMessage(msg.Id) != nil {
		return fmt.Errorf("message with id %d already exists", msg.Id)
	}

	inbox.Messages[msg.Id] = msg

	return nil
}

func EmptyInbox() Inbox {
	return Inbox{
		Messages: map[uint64]models.Email{
			1337: {Id: 1337, Sender: "hello@veiled.io", SenderName: "Veiled", Subject: "Received emails will appear here.", Html: "Send a message to the Veiled address, and it will show up here. Try it out!"},
		},
	}
}
