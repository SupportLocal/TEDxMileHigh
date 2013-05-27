package main

import (
	"fmt"
	"net/http"
	"os"

	"labix.org/v2/mgo"

	"supportlocal/TEDxMileHigh/handlers/current_message"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"supportlocal/TEDxMileHigh/lib/json"
	"supportlocal/TEDxMileHigh/mongo"
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

	http.Handle("/currentMessage", eventsource)
	http.Handle("/", router.New())

	go func() {

		session, err := mgo.Dial("localhost")
		fatal.If(err)

		database := session.DB("tedx")
		messages := mongo.NewMessagesRepo(database)

		fatal.If(messages.Tail(func(message mongo.Message) {
			data := fmt.Sprintf("%s", json.MustMarshal(message))
			eventsource.SendMessage(data, "", message.Id.String())
		}))

	}()

	go func() {
		if pidFile, err := os.Create("./TEDxMileHigh.pid"); err == nil {
			pidFile.WriteString(fmt.Sprintf("%v", os.Getpid()))
			pidFile.Close()
		}
	}()

	fatal.If(http.ListenAndServe(":9000", nil))
}
