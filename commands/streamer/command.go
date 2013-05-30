package streamer

import (
	"log"
	"time"

	"github.com/darkhelmet/twitterstream"
	"github.com/laurent22/toml-go/toml"
	"labix.org/v2/mgo"

	"supportlocal/TEDxMileHigh/commands"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"supportlocal/TEDxMileHigh/mongo"
)

func init() { commands.Register(command{"streamer"}) }

type command struct{ name string }

func (cmd command) Name() string           { return cmd.name }
func (cmd command) CanCreatePidFile() bool { return true }

func (cmd command) Run(config toml.Document) {

	var (
		debug    = config.GetBool("debug") || config.GetBool("streamer.debug")
		track    = config.GetString("streamer.track", "supportlocal")
		username = config.GetString("twitter.username")
		password = config.GetString("twitter.password")

		client = twitterstream.NewClient(username, password)
	)

	session, err := mgo.Dial(config.GetString("mongo.dial"))
	fatal.If(err)

	mongo.Database = session.DB(config.GetString("mongo.database"))
	inboundMessageRepo := mongo.InboundMessageRepo()
	twitterCrosswalkRepo := mongo.TwitterCrosswalkRepo()

	decode := func(conn *twitterstream.Connection) {
		for {
			tweet, err := conn.Next()
			if err != nil {
				log.Printf("streamer: decoding failed: %s", err)
				return
			}

			if debug {
				log.Printf("streamer: %s said: %s", tweet.User.ScreenName, tweet.Text)
			}

			// save it as an inbound message

			twitterCrosswalk, err := twitterCrosswalkRepo.FindOrCreate(tweet.Id)
			fatal.If(err)

			fatal.If(inboundMessageRepo.Save(&mongo.InboundMessage{
				Id:      twitterCrosswalk.InternalId,
				Comment: tweet.Text,
				Name:    "@" + tweet.User.ScreenName,
			}))
		}
	}

	for {
		conn, err := client.Track(track)
		if err != nil {
			log.Println("streamer: tracking failed, sleeping")
			time.Sleep(30 * time.Second)
			continue
		}
		decode(conn)
	}
}
