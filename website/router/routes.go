package router

import (
	"net/http"
	"supportlocal/TEDxMileHigh/website/handlers"
	"supportlocal/TEDxMileHigh/website/handlers/form"
	"supportlocal/TEDxMileHigh/website/handlers/jumbotron"
)

type route struct {
	name      string
	methods   string
	path      string
	adminOnly bool
	handler   func(http.ResponseWriter, *http.Request)
}

var routes = []route{

	{
		handler: form.Get,
		methods: "GET",
		name:    "get-form",
		path:    "/",
	},

	{
		handler:   handlers.AdminHome,
		adminOnly: true,
		methods:   "GET",
		name:      "admin-home",
		path:      "/admin",
	},

	{
		handler:   handlers.BlockMessages,
		adminOnly: true,
		methods:   "POST",
		name:      "block-messages",
		path:      "/admin/messages/block",
	},

	{
		handler: form.Post,
		methods: "POST",
		name:    "post-form",
		path:    "/",
	},

	{
		handler: jumbotron.Get,
		methods: "GET",
		name:    "jumbotron",
		path:    "/jumbotron",
	},
}
