package httputil

import (
	"net/http"
	"supportlocal/TEDxMileHigh/lib/fatal"
)

func MustWriteJson(w http.ResponseWriter, v interface{}) {
	fatal.If(WriteJson(w, v))
}
