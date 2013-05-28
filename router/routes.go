package router

import (
	"net/http"
	"supportlocal/TEDxMileHigh/handlers"
	"supportlocal/TEDxMileHigh/handlers/form"
)

type route struct {
	name,
	methods,
	path string
	tlsOnly bool
	handler func(http.ResponseWriter, *http.Request)
}

var routes = []route{

	{
		handler: form.Get,
		methods: "GET",
		name:    "get-form",
		path:    "/",
	},

	{
		//handler: ??,
		methods: "GET",
		name:    "admin-home",
		path:    "/admin",
	},

	{
		handler: form.Post,
		methods: "POST",
		name:    "post-form",
		path:    "/",
	},

	{
		handler: handlers.Jumbotron,
		methods: "GET",
		name:    "jumbotron",
		path:    "/jumbotron",
	},
}
