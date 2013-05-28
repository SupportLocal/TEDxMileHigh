package form

import (
	"strings"
	"supportlocal/TEDxMileHigh/mongo"
)

type form struct {
	Comment string            `json:"comment"`
	Email   string            `json:"email"`
	Name    string            `json:"name"`
	Errors  map[string]string `json:"errors"`
}

func (f *form) valid() bool {
	f.Errors = make(map[string]string)

	if email := strings.TrimSpace(f.Email); len(email) == 0 {
		f.Errors["email"] = "is required"
	}

	if name := strings.TrimSpace(f.Name); len(name) == 0 {
		f.Errors["name"] = "is required"
	}

	if comment := strings.TrimSpace(f.Comment); len(comment) == 0 {
		f.Errors["comment"] = "is required"
	}

	return len(f.Errors) == 0
}

func (f form) toInBoundMessage() mongo.InboundMessage {
	return mongo.InboundMessage{
		Comment: f.Comment,
		Email:   f.Email,
		Name:    f.Name,
	}
}
