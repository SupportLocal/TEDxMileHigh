package handlers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"supportlocal/TEDxMileHigh/domain/models"
	"supportlocal/TEDxMileHigh/lib/pager"
	"supportlocal/TEDxMileHigh/redis"
	"supportlocal/TEDxMileHigh/website/layout"
)

func AdminHome(w http.ResponseWriter, r *http.Request) {
	var (
		messageRepo = redis.MessageRepo()

		query = r.URL.Query()
		pager = pager.Parse(query)

		current  models.Message
		messages models.Messages
		err      error

		tail bytes.Buffer
	)

	if current, err = messageRepo.Tail(); err != nil {
		log.Printf("website: handlers.AdminHome messageRepo.Tail failed %q", err)
	}

	if messages, err = messageRepo.Paginate(pager); err != nil {
		log.Printf("website: handlers.AdminHome messageRepo.Paginate failed %q", err)
	}

	tail.WriteString(scriptIsland("data-pool", struct {
		Messages models.Messages `json:"messages"`
		Current  models.Message  `json:"current"`
	}{
		Messages: messages,
		Current:  current,
	}))

	tail.WriteString(`<script src="/js/admin_home.js"></script>`)

	mustWriteHtml(w, layout.DefaultLayout{
		Title: "SupportLocal | TEDxMilehigh",
		Tail:  template.HTML(tail.String()),
	})
}
