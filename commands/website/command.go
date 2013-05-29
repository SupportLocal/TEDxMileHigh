package website

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

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

func (cmd command) Run(args []string, config toml.Document) {

	session, err := mgo.Dial(config.GetString("mongo.dial"))
	fatal.If(err)
	mongo.Database = session.DB(config.GetString("mongo.database"))

	eventsource := es.New(nil)

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

	go func() { // mongo cursor; sends messages to eventsource
		currentMessageRepo := mongo.CurrentMessageRepo()

		fatal.If(currentMessageRepo.Tail(func(msg mongo.CurrentMessage) {
			data := fmt.Sprintf("%s", json.MustMarshal(msg))
			eventsource.SendMessage(data, "", msg.Id.Hex())
			log.Printf("website: /currentMessage sent to %d consumers", eventsource.ConsumersCount())
		}))
	}()

	<-make(chan bool) // don't exit
}
