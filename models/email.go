package models

type Email struct {
	Id         uint64   `json:"id"`
	Sender     string   `json:"sender"`
	Recipients []string `json:"recipients"`
	SenderName string   `json:"sender_name"`
	Subject    string   `json:"subject"`
	Text       string   `json:"text"`
	Html       string   `json:"html"`
	ReceivedAt string   `json:"received_at"`
}
