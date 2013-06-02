package form

import (
	"log"
	"net/http"
	"supportlocal/TEDxMileHigh/redis"
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
		messageRepo := redis.MessageRepo()
		message := form.toMessage()
		if err = messageRepo.Save(&message); err != nil {
			log.Printf("website: form.Post save failed %q", err)
		}
	}

	mustWriteJson(w, form)
}
