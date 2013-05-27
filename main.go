package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"supportlocal/TEDxMileHigh/handlers/current_message"
	"supportlocal/TEDxMileHigh/lib/json"
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

	eventsource := current_message.NewEventSource()

	type message struct {
		Id      int    `json:"id"`
		Author  string `json:"author"`
		Comment string `json:"comment"`
	}

	go func() {
		id := 1
		for {

			data := fmt.Sprintf("%s", json.MustMarshal(message{
				Id:      id,
				Author:  fmt.Sprintf("@foo %d", id),
				Comment: fmt.Sprintf("dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna ali. %d", id),
			}))

			eventsource.SendMessage(data, "", strconv.Itoa(id))
			id++
			time.Sleep(5 * time.Second)
		}
	}()

	http.Handle("/currentMessage", eventsource)
	http.Handle("/", router.New())

	go func() {
		if pidFile, err := os.Create("./TEDxMileHigh.pid"); err == nil {
			pidFile.WriteString(fmt.Sprintf("%v", os.Getpid()))
			pidFile.Close()
		}
	}()

	http.ListenAndServe(":9000", nil)
}
