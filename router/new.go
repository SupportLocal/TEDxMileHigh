package router

import (
	"github.com/gorilla/mux"
)

var (
	router       = mux.NewRouter().StrictSlash(true)
	routesLoaded = false
)

func New() *mux.Router {
	if routesLoaded {
		return router
	}

	for _, r := range routes {
		hf := r.handler

		router.
			Path(r.path).
			Methods(r.methods).
			HandlerFunc(hf).
			Name(r.name)
	}

	routesLoaded = true

	return router
}
