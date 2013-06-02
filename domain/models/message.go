package models

type (
	Messages []Message
	Message  struct {
		Id      int    `json:"id"`
		Author  string `json:"author"`
		Comment string `json:"comment"`
		Email   string `json:"email"`
	}
)
