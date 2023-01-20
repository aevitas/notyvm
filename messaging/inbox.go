package messaging

import "aevitas.dev/veiled/models"

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
