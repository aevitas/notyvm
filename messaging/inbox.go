package messaging

import (
	"fmt"

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
			1337: {Id: 1337, Sender: "hello@veiled.io", SenderName: "Veiled", Subject: "Received emails will appear here.", Text: "Send a message to the Veiled address, and they'll show up here. Try it out!"},
		},
	}
}
