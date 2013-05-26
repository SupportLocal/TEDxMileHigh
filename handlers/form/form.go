package form

import (
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {}

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
		<div class="container">
			<div class="form">
				<form method="POST">

					<label for="email">Email Address</label>
					<input id="email" name="email" type="text" placeholder="Please enter your email address">

					<label for="name">Your Name</label>
					<input id="name" name="name" type="text" placeholder="Please enter your name">

					<label for="comment">What does Support Local mean to you?</label>
					<textarea id="comment" name="comment" placeholder="Type something&hellip;">@SupportLocal means </textarea>

					<button type="submit">Submit</button>
				</form>
			</div>
		</div>
		<div id="tail" style="position: absolute; top: -10000px; height: 0px; width: 0px;">
			<script src="/js/form.js"></script>
		</div>
	</body>
</html>
`
