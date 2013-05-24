package handlers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(homeHtml))
}

const homeHtml = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
		<title>SupportLocal | TEDxMilehigh</title>

		<link href="/css/screen.css" rel="stylesheet" type="text/css" />

		<script src="http://localhost:35729/livereload.js"></script>

	</head>
	<body>
		<div class="floater"></div>
		<div class="container">
			<div class="center">
				<div class="left">
					<img src="/img/headlineimg.png" width="647" height="402" />
					<p>Send a tweet to @SupportLocal with your answer or visit tedx.supportlocal.com</p>
				</div>
				<div class="right">
					<p>@Supportlocal Lorem ipsum dolor sit amet, <strong>consectetuer</strong> adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna ali.</p>
					<span class="from">@benogren</span>
				</div>
			</div>
		</div>
		<div id="tail" style="position: absolute; top: -10000px; height: 0px; width: 0px;">
			<script src="/js/home.js"></script>
		</div>
	</body>
</html>
`
