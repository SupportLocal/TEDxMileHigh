package handlers

import (
	"log"
	"net/http"
	"supportlocal/TEDxMileHigh/domain/models"
	"supportlocal/TEDxMileHigh/lib/pager"
	"supportlocal/TEDxMileHigh/redis"
)

func AdminHome(w http.ResponseWriter, r *http.Request) {
	var (
		messageRepo = redis.MessageRepo()

		query = r.URL.Query()
		pager = pager.Parse(query)

		messages models.Messages
		err      error
	)

	if messages, err = messageRepo.Paginate(pager); err != nil {
		log.Printf("website: handlers.AdminHome messageRepo.Paginate failed %q", err)
	}

	_ = messages // deleteme

	//	mustWriteHtml(w, view{
	//		Comment: template.HTML(message.Comment),
	//		Author:  message.Author,
	//	})

	w.Write([]byte(adminHomeHtml))
}

const adminHomeHtml = `<!DOCTYPE html>
<html>
	<head>
		<title>SupportLocal | TEDxMilehigh</title>

		<link href="/css/screen.css"     media="screen, projection" rel="stylesheet" type="text/css" />
		<link href="/vendor/animate.css" media="screen, projection" rel="stylesheet" type="text/css" />

		<script src="http://localhost:35729/livereload.js"></script>

	</head>
	<body>
		<div class="floater"></div>
		<div class="container"></div>
		<div id="tail" style="position: absolute; top: -10000px; height: 0px; width: 0px;">
			<script src="/vendor/jquery-2.0.1.js"></script>
			<script src="/vendor/can.jquery-1.1.5.js"></script>
			<script src="/js/admin_home.js"></script>
		</div>
	</body>
</html>
`
