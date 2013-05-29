package httputil

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
)

type htmlRenderer interface {
	WriteHtmlTo(io.Writer) error
}

func WriteHtml(w http.ResponseWriter, renderers ...htmlRenderer) (err error) {
	var htmlBuffer bytes.Buffer

	for _, renderer := range renderers {
		if err = renderer.WriteHtmlTo(io.Writer(&htmlBuffer)); err != nil {
			return
		}
	}

	w.Header().Set("Content-Length", strconv.Itoa(htmlBuffer.Len()))
	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlBuffer.Bytes())

	return
}
