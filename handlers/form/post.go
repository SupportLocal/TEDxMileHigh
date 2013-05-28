package form

import (
	"log"
	"net/http"
	"supportlocal/TEDxMileHigh/mongo"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var (
		form form
		err  error
	)

	if err = readJson(r, &form); err != nil {
		log.Printf("readJson failed %q", err)
	}

	if form.valid() { // save it
		inboundMessage := form.toInBoundMessage()
		inboundMessageRepo := mongo.InboundMessageRepo()
		if err = inboundMessageRepo.Save(&inboundMessage); err != nil {
			log.Printf("inboundMessageRepo.Save failed %q", err)
		}
	}

	mustWriteJson(w, form)
}
