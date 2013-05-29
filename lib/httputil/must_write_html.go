package httputil

import (
	"net/http"
	"supportlocal/TEDxMileHigh/lib/fatal"
)

func MustWriteHtml(w http.ResponseWriter, renderers ...htmlRenderer) {
	fatal.If(WriteHtml(w, renderers...))
}
