package form

import (
	"net/http"
	"supportlocal/TEDxMileHigh/mongo"
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

		inboundMessage := form.toInBoundMessage()
		inboundMessageRepo := mongo.InboundMessageRepo()
		if err = inboundMessageRepo.Save(&inboundMessage); err != nil {
			// todo log it
		}
	}

	mustWriteJson(w, form)
}
