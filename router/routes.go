package router

import (
	"net/http"
	"supportlocal/TEDxMileHigh/handlers"
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
		handler: handlers.Form,
		methods: "GET",
		name:    "form",
		path:    "/form",
	},
}
