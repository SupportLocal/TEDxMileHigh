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

		pendingPager = pager.Parse(query)
		blockedPager = pager.Parse(query)

		current models.Message
		pending models.Messages
		blocked models.Messages
		err     error

		tail bytes.Buffer
	)

	if current, err = messageRepo.Tail(); err != nil {
		log.Printf("website: handlers.AdminHome messageRepo.Tail failed %q", err)
	}

	if pending, err = messageRepo.PaginatePending(pendingPager); err != nil {
		log.Printf("website: handlers.AdminHome messageRepo.PaginatePending failed %q", err)
	}

	if blocked, err = messageRepo.PaginateBlocked(blockedPager); err != nil {
		log.Printf("website: handlers.AdminHome messageRepo.PaginateBlocked failed %q", err)
	}

	tail.WriteString(scriptIsland("data-pool", struct {
		Current      models.Message  `json:"current"`
		Pending      models.Messages `json:"pending"`
		PendingPager pager.Pager     `json:"pendingPager"`
		Blocked      models.Messages `json:"blocked"`
		BlockedPager pager.Pager     `json:"blockedPager"`
	}{
		Current:      current,
		Pending:      pending,
		PendingPager: pendingPager,
		Blocked:      blocked,
		BlockedPager: blockedPager,
	}))

	tail.WriteString(`<script src="/js/admin_home.js"></script>`)

	mustWriteHtml(w, layout.DefaultLayout{
		Title: "SupportLocal | TEDxMilehigh",
		Tail:  template.HTML(tail.String()),
	})
}
