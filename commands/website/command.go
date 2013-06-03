package website

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	es "github.com/antage/eventsource/http"
	"github.com/laurent22/toml-go/toml"

	"supportlocal/TEDxMileHigh/commands"
	"supportlocal/TEDxMileHigh/domain/pubsub"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"supportlocal/TEDxMileHigh/lib/json"
	"supportlocal/TEDxMileHigh/redis"
	"supportlocal/TEDxMileHigh/website/router"
)

func init() { commands.Register(command{"website"}) }

type command struct{ name string }

func (cmd command) Name() string           { return cmd.name }
func (cmd command) CanCreatePidFile() bool { return true }

func (cmd command) Run(config toml.Document) {
	debug := config.GetBool("debug") || config.GetBool("website.debug")

	eventsource := es.New(nil)

	go func() {

		messageRepo := redis.MessageRepo()

		subscription := messageRepo.Subscribe(pubsub.MessageBlocked, pubsub.MessageCycled)
		defer subscription.Unsubscribe()

		for {
			channel, message, err := subscription.Receive()
			fatal.If(err)

			if debug {
				log.Printf("website: %s %d", channel, message.Id)
			}

			data := fmt.Sprintf("%s", json.MustMarshal(struct {
				Id      int    `json:"id"`
				Author  string `json:"author"`
				Comment string `json:"comment"`
			}{
				Id:      message.Id,
				Author:  message.Author,
				Comment: message.Comment,
			}))

			eventsource.SendMessage(data, channel.String(), "")
		}
	}()

	go func() { // http dance
		for _, assetPath := range []string{"/css/", "/ejs/", "/img/", "/js/", "/vendor/"} {
			pt := filepath.Join(config.GetString("website.assets"), assetPath)
			fs := http.FileServer(http.Dir(pt))
			http.Handle(assetPath, http.StripPrefix(assetPath, fs))
		}

		http.Handle("/message/events", eventsource)
		http.Handle("/", router.New(config))
		fatal.If(http.ListenAndServe(config.GetString("website.listen-addr"), nil))
	}()

	<-make(chan bool) // don't exit
}
