package form

import (
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var (
		form = &form{}
		err  error
	)

	if err = readJson(r, form); err != nil {
		// todo log it
	}

	if form.valid() {
		// save it
	}

	mustWriteJson(w, form)
}
