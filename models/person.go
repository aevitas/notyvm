package models

type Person struct {
	Seed         int    `json:"seed"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
}
