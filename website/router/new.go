package router

import (
	"encoding/base64"
	"github.com/gorilla/mux"
	"github.com/laurent22/toml-go/toml"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

var (
	router         = mux.NewRouter().StrictSlash(true)
	routesLoaded   = false
	administrators map[string]string
)

func New(config toml.Document) *mux.Router {
	if routesLoaded {
		return router
	}

	administrators = make(map[string]string)
	for _, administrator := range config.GetArray("website.administrators") {
		auth := strings.SplitN(administrator.AsString(), ":", 2)
		if len(auth) == 2 {
			username, password := auth[0], auth[1]
			administrators[username] = password
		}
	}

	for _, r := range routes {

		hf := timingFilter(r.name, r.handler)

		if r.adminOnly {
			hf = requireAdmin(hf)
		}

		router.
			Path(r.path).
			Methods(r.methods).
			HandlerFunc(hf).
			Name(r.name)
	}

	routesLoaded = true

	return router
}

func timingFilter(name string, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			if x := recover(); x == nil {
				log.Printf("%s %s | %s %s", r.Method, r.RequestURI, name, time.Since(start))
			} else {
				log.Printf("%s %s | %s %s | error: %v\n%s", r.Method, r.RequestURI, name, time.Since(start), x, debug.Stack())
			}
		}(time.Now())

		fn(w, r)
	}
}

func requireAdmin(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		un, up := getBasicAuthorization(r)
		sp, ok := administrators[un]
		authorized := (ok && up == sp)

		if !authorized {
			w.Header().Set("WWW-Authenticate", `Basic realm="SupportLocal | TEDxMileHigh"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			fn(w, r)
		}
	}
}

func getBasicAuthorization(r *http.Request) (username, password string) {
	authorization := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(authorization) == 2 {
		method := authorization[0]
		secret := authorization[1]
		switch method {
		case "Basic":
			if dec, err := base64.StdEncoding.DecodeString(secret); err == nil {
				pair := strings.SplitN(string(dec), ":", 2)
				if len(pair) == 2 {
					username = pair[0]
					password = pair[1]
				}
			}
		}
	}
	return
}
