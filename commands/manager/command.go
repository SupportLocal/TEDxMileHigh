package manager

import (
	"time"

	"github.com/laurent22/toml-go/toml"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"supportlocal/TEDxMileHigh/commands"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"supportlocal/TEDxMileHigh/mongo"
)

func init() { commands.Register(command{"manager"}) }

type command struct{ name string }

func (cmd command) Name() string           { return cmd.name }
func (cmd command) CanCreatePidFile() bool { return true }

// periodically copies inbound messages in to the current message collection
func (cmd command) Run(args []string, config toml.Document) {

	session, err := mgo.Dial(config.GetString("mongo.dial"))
	fatal.If(err)
	mongo.Database = session.DB(config.GetString("mongo.database"))

	currentMessageRepo := mongo.CurrentMessageRepo()
	inboundMessageRepo := mongo.InboundMessageRepo()

	duration, err := time.ParseDuration(config.GetString("manager.ticker-duration"))
	fatal.If(err)
	ticker := time.NewTicker(duration)

	go func() {

		for _ = range ticker.C {

			currentMessage, err := currentMessageRepo.Last()
			if err != nil && err != mgo.ErrNotFound {
				panic(err)
			}

			lastId := currentMessage.Id
			if !lastId.Valid() {
				epoch := time.Unix(0, 0)
				lastId = bson.NewObjectIdWithTime(epoch)
			}

			inboundMessage, err := inboundMessageRepo.Next(lastId)
			if err != nil && err != mgo.ErrNotFound {
				panic(err)
			}

			if inboundMessage.Valid() { // create a new current message
				currentMessage = inboundMessage.ToCurrentMessage()
				fatal.If(currentMessageRepo.Save(&currentMessage))
				fatal.If(inboundMessageRepo.Converted(inboundMessage.Id))
			}

		}
	}()

	<-make(chan bool) // don't exit
}
