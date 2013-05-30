package website

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	es "github.com/antage/eventsource/http"
	"github.com/laurent22/toml-go/toml"
	"labix.org/v2/mgo"

	"supportlocal/TEDxMileHigh/commands"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"supportlocal/TEDxMileHigh/lib/json"
	"supportlocal/TEDxMileHigh/mongo"
	"supportlocal/TEDxMileHigh/router"
)

func init() { commands.Register(command{"website"}) }

type command struct{ name string }

func (cmd command) Name() string           { return cmd.name }
func (cmd command) CanCreatePidFile() bool { return true }

func (cmd command) Run(config toml.Document) {
	debug := config.GetBool("debug") || config.GetBool("website.debug")
	_ = debug

	session, err := mgo.Dial(config.GetString("mongo.dial"))
	fatal.If(err)
	mongo.Database = session.DB(config.GetString("mongo.database"))

	eventsource := es.New(nil)

	go func() {
		duration, err := time.ParseDuration(config.GetString("website.heartbeat"))
		fatal.If(err)
		ticker := time.NewTicker(duration)

		currentMessageRepo := mongo.CurrentMessageRepo()

		for _ = range ticker.C {
			currentMessage, err := currentMessageRepo.Last()
			if err != nil && err != mgo.ErrNotFound {
				panic(err)
			}

			data := fmt.Sprintf("%s", json.MustMarshal(currentMessage))
			eventsource.SendMessage(data, "", currentMessage.Id.Hex())
		}
	}()

	go func() {
		duration, err := time.ParseDuration(config.GetString("website.heartbeat"))
		fatal.If(err)
		ticker := time.NewTicker(duration)

		for _ = range ticker.C {
			log.Printf("website: /currentMessage consumers: %d", eventsource.ConsumersCount())
		}
	}()

	go func() { // http dance
		for _, assetPath := range []string{"/css/", "/ejs/", "/img/", "/js/", "/vendor/"} {
			pt := filepath.Join(config.GetString("website.assets"), assetPath)
			fs := http.FileServer(http.Dir(pt))
			http.Handle(assetPath, http.StripPrefix(assetPath, fs))
		}

		http.Handle("/currentMessage", eventsource)
		http.Handle("/", router.New(config))
		fatal.If(http.ListenAndServe(config.GetString("website.listen-addr"), nil))
	}()

	/*
	go func() { // mongo cursor; sends messages to eventsource
		currentMessageRepo := mongo.CurrentMessageRepo()

		fatal.If(currentMessageRepo.Tail(func(msg mongo.CurrentMessage) {
			log.Printf("website: currentMessageRepo.Tail called %q", msg.Id)

			data := fmt.Sprintf("%s", json.MustMarshal(msg))
			eventsource.SendMessage(data, "", msg.Id.Hex())

			if debug {
				log.Printf("website: /currentMessage sent to %d consumers", eventsource.ConsumersCount())
			}
		}))
	}()
	*/

	<-make(chan bool) // don't exit
}
