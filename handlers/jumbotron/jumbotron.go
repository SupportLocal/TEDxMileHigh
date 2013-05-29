package jumbotron

import (
	"fmt"
	"html/template"
	"io"
	"labix.org/v2/mgo"
	"net/http"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"supportlocal/TEDxMileHigh/lib/httputil"
	"supportlocal/TEDxMileHigh/mongo"
)

var (
	viewTemplate  *template.Template
	mustWriteHtml = httputil.MustWriteHtml
)

func init() {
	t, e := template.New("jumbotron/viewHtml").Parse(viewHtml)
	fatal.If(e)
	viewTemplate = t
}

func Get(w http.ResponseWriter, r *http.Request) {
	currentMessageRepo := mongo.CurrentMessageRepo()

	currentMessage, err := currentMessageRepo.Last()
	if err != nil && err != mgo.ErrNotFound {
		panic(err)
	}

	panic(fmt.Errorf("gah!!"))

	mustWriteHtml(w, view{
		Comment: template.HTML(currentMessage.Comment),
		Author:  currentMessage.Author,
	})
}

type view struct {
	Comment template.HTML
	Author  string
}

func (v view) WriteHtmlTo(w io.Writer) error {
	return viewTemplate.Execute(w, v)
}

const viewHtml = `<!DOCTYPE html>
<html>
	<head>
		<title>SupportLocal | TEDxMilehigh</title>

		<link href="/css/screen.css"     media="screen, projection" rel="stylesheet" type="text/css" />
		<link href="/vendor/animate.css" media="screen, projection" rel="stylesheet" type="text/css" />

		<script src="http://localhost:35729/livereload.js"></script>

	</head>
	<body id="jumbotron">
		<div class="floater"></div>
		<div class="container">
			<div class="center">
				<div class="left">
					<img src="/img/headlineimg.png" width="647" height="402" />
					<p>Send a tweet to @SupportLocal with your answer or visit tedx.supportlocal.com</p>
				</div>
				<div id="flipboard" class="right">
					<p class="comment animated">{{ .Comment }}</p>
					<span class="author animated">{{ .Author }}</span>
				</div>
			</div>
		</div>
		<div id="tail" style="position: absolute; top: -10000px; height: 0px; width: 0px;">
			<script src="/vendor/jquery-2.0.1.js"></script>
			<script src="/vendor/can.jquery-1.1.5.js"></script>
			<script src="/js/jumbotron.js"></script>
		</div>
	</body>
</html>
`
