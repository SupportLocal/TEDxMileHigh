package layout

import (
	"bytes"
	"html/template"
	"io"
	"supportlocal/TEDxMileHigh/lib/fatal"
)

type htmlRenderer interface {
	WriteHtmlTo(io.Writer) error
}

type DefaultLayout struct {
	Title string
	Head  template.HTML
	View  htmlRenderer
	Tail  template.HTML
}

func (dl DefaultLayout) Main() template.HTML {
	var buffer bytes.Buffer

	if dl.View != nil {
		if err := dl.View.WriteHtmlTo(&buffer); err != nil {
			panic(err) // shouldn't happen, complain loudly if it does
		}
	}

	return template.HTML(buffer.String())
}

func (dl DefaultLayout) LiveRelaodHook() template.HTML {
	// TODO conditionalize this; based on config
	return template.HTML(`<script src="http://localhost:35729/livereload.js"></script>`)
}

func (dl DefaultLayout) WriteHtmlTo(w io.Writer) error {
	return defaultLayoutTemplate.Execute(w, dl)
}

var defaultLayoutTemplate *template.Template

func init() {
	t, e := template.New("layout/DefaultLayout").Parse(defaultLayoutHtml)
	fatal.If(e)
	defaultLayoutTemplate = t
}

const defaultLayoutHtml = `<!DOCTYPE html>
<html>
	<head>
		<title>{{ .Title }}</title>
		<link href="/vendor/animate.css" media="screen, projection" rel="stylesheet" type="text/css" />
		<link href="/css/screen.css"     media="screen, projection" rel="stylesheet" type="text/css" />
		{{ .Head }}
	</head>
	<body>
		{{ .Main }}
		<div id="tail" style="position: absolute; top: -10000px; height: 0px; width: 0px;">
			<script src="/vendor/jquery-2.0.1.js"></script>
			<script src="/vendor/can.jquery-1.1.5.js"></script>
			{{ .Tail }}
			{{ .LiveRelaodHook }}
		</div>
	</body>
</html>
`
