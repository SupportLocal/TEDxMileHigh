package form

type form struct {
	Email   string            `json:"email"`
	Name    string            `json:"name"`
	Comment string            `json:"comment"`
	Errors  map[string]string `json:"errors"`
}

func (f *form) valid() bool {
	f.Errors = make(map[string]string)

	f.Errors["email"] = "is required"
	f.Errors["name"] = "is required"
	f.Errors["comment"] = "is required"

	return len(f.Errors) == 0
}
