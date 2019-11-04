package models

type User struct {
	ID                 int64  `json:"id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	Bearer             string `json:"bearer"`
	Passwordhash       string `json:"password"`
	EmailVerified      bool   `json:"emailVerified"`
	EmailVerifiedToken string `json:"emailVerifiedToken"`
}
