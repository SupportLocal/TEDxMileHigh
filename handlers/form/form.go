package form

import (
	"net/http"
)

type form struct {
	Email   string            `json:"email"`
	Name    string            `json:"name"`
	Comment string            `json:"comment"`
	Errors  map[string]string `json:"errors"`
}

func (f *form) valid() bool {
	f.Errors = make(map[string]string)

	f.Errors["email"] = "is required"
	f.Errors["name"] = "is required"
	f.Errors["comment"] = "is required"

	return len(f.Errors) == 0
}

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
	}

	mustWriteJson(w, form)
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(formHtml))
}

const formHtml = `<!DOCTYPE html>
<html>
	<head>
		<title>SupportLocal | TEDxMilehigh</title>
		<link href="/css/screen.css" media="screen, projection" rel="stylesheet" type="text/css" />
		<script src="http://localhost:35729/livereload.js"></script>
	</head>
	<body>
		<div class="floater"></div>
		<div id="container" class="container"></div>
		<div id="tail" style="position: absolute; top: -10000px; height: 0px; width: 0px;">
			<script src="/vendor/jquery-2.0.1.js"></script>
			<script src="/vendor/can.jquery-1.1.5.js"></script>
			<script src="/js/form.js"></script>
		</div>
	</body>
</html>
`
