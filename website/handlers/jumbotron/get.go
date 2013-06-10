package jumbotron

import (
	"html/template"
	"log"
	"net/http"
	"supportlocal/TEDxMileHigh/redis"
	"supportlocal/TEDxMileHigh/website/layout"
)

func Get(w http.ResponseWriter, r *http.Request) {
	messageRepo := redis.MessageRepo()

	message, err := messageRepo.Tail()
	if err != nil {
		log.Printf("website: jumbotron.Get messageRepo.Tail failed %q", err)
	}

	mustWriteHtml(w, layout.DefaultLayout{
		Title:  "SupportLocal | TEDxMilehigh",
		BodyId: "jumbotron",
		Tail:   template.HTML(tail),
		View: view{
			Comment: template.HTML(message.Comment),
			Author:  message.Author,
		},
	})
}
