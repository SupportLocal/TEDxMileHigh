package streamer

import (
	"github.com/darkhelmet/twitterstream"
	"github.com/laurent22/toml-go/toml"
	"log"
	"time"

	"supportlocal/TEDxMileHigh/commands"
	"supportlocal/TEDxMileHigh/domain/models"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"supportlocal/TEDxMileHigh/redis"
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

		messageRepo = redis.MessageRepo()
		crosswalk   = redis.TwitterCrosswalk(messageRepo)
	)

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

			messageId, err := crosswalk.MessageIdFor(tweet.Id)
			fatal.If(err)

			fatal.If(messageRepo.Save(&models.Message{
				Id:      messageId,
				Comment: tweet.Text,
				Author:  "@" + tweet.User.ScreenName,
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
