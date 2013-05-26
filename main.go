package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"net/http"
	"os"
	"supportlocal/TEDxMileHigh/router"
)

func main() {
	// TODO "./assets"           should come from config .. or command line args --assets=
	// TODO "./TEDxMileHigh.pid" should come from config .. or command line args --pid-file=

	for _, assetPath := range []string{"/css/", "/img/", "/js/", "/vendor/"} {
		dr := http.Dir("./assets" + assetPath)
		fs := http.FileServer(dr)
		http.Handle(assetPath, http.StripPrefix(assetPath, fs))
	}

	http.Handle("/echo", websocket.Handler(func(ws *websocket.Conn) {
		io.Copy(ws, ws)
	}))

	http.Handle("/", router.New())

	go func() {
		if pidFile, err := os.Create("./TEDxMileHigh.pid"); err == nil {
			pidFile.WriteString(fmt.Sprintf("%v", os.Getpid()))
			pidFile.Close()
		}
	}()

	http.ListenAndServe(":9000", nil)
}
