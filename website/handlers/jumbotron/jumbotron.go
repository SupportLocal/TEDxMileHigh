package jumbotron

import (
	"html/template"
	"io"
	"supportlocal/TEDxMileHigh/lib/fatal"
)

var viewTemplate *template.Template

func init() {
	t, e := template.New("jumbotron/viewHtml").Parse(viewHtml)
	fatal.If(e)
	viewTemplate = t
}

type view struct {
	Comment template.HTML
	Author  string
}

func (v view) WriteHtmlTo(w io.Writer) error {
	return viewTemplate.Execute(w, v)
}

const tail = `<script src="/js/jumbotron.js"></script>`

const viewHtml = `
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
`
