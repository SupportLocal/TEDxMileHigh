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
		handler: handlers.Home,
		methods: "GET",
		name:    "home",
		path:    "/",
	},
	{
		handler: form.Get,
		methods: "GET",
		name:    "get-form",
		path:    "/form",
	},
	{
		handler: form.Post,
		methods: "POST",
		name:    "post-form",
		path:    "/form",
	},
}
