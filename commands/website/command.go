package website

import (
	"fmt"
	"net/http"
	"os"

	es "github.com/antage/eventsource/http"
	"labix.org/v2/mgo"

	"supportlocal/TEDxMileHigh/commands"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"supportlocal/TEDxMileHigh/lib/json"
	"supportlocal/TEDxMileHigh/mongo"
	"supportlocal/TEDxMileHigh/router"
)

func init() { commands.Register(command{"website"}) }

type command struct{ name string }

func (cmd command) Name() string { return cmd.name }

func (cmd command) Run(args []string) {

	// TODO "./assets"                   should come from config .. or command line args --assets=
	// TODO "./TEDxMileHigh-website.pid" should come from config .. or command line args --pid-file=
	// TODO "localhost" and "tedx"       should come from config .. or command line args --mgo-dial= and --mgo-db=

	eventsource := es.New(nil)

	go func() { // http dance
		for _, assetPath := range []string{"/css/", "/img/", "/js/", "/vendor/"} {
			dr := http.Dir("./assets" + assetPath)
			fs := http.FileServer(dr)
			http.Handle(assetPath, http.StripPrefix(assetPath, fs))
		}

		http.Handle("/currentMessage", eventsource)
		http.Handle("/", router.New())
		fatal.If(http.ListenAndServe(":9000", nil))
	}()

	go func() { // mongo dance
		session, err := mgo.Dial("localhost")
		fatal.If(err)

		currentMessageRepo := mongo.CurrentMessageRepo(session.DB("tedx"))

		fatal.If(currentMessageRepo.Tail(func(message mongo.Message) {
			data := fmt.Sprintf("%s", json.MustMarshal(message))
			eventsource.SendMessage(data, "", message.Id.String())
		}))
	}()

	go func() { // create our pid file
		if pidFile, err := os.Create("./TEDxMileHigh-website.pid"); err == nil {
			pidFile.WriteString(fmt.Sprintf("%v", os.Getpid()))
			pidFile.Close()
		}
	}()

	<-make(chan bool) // don't exit
}
