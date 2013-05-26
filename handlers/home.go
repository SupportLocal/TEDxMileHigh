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
		<title>SupportLocal | TEDxMilehigh</title>

		<link href="/css/screen.css"     media="screen, projection" rel="stylesheet" type="text/css" />
		<link href="/vendor/animate.css" media="screen, projection" rel="stylesheet" type="text/css" />

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
				<div id="flipboard" class="right">
					<p class="comment animated">
						@Supportlocal Lorem ipsum dolor sit amet, <strong>consectetuer</strong>
						adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet
						dolore magna ali.
					</p>
					<span class="author animated">@benogren</span>
				</div>
			</div>
		</div>
		<div id="tail" style="position: absolute; top: -10000px; height: 0px; width: 0px;">
			<script src="/vendor/jquery-2.0.1.js"></script>
			<script src="/vendor/can.jquery-1.1.5.js"></script>
			<script src="/js/home.js"></script>
		</div>
	</body>
</html>
`
