package manager

import (
	"log"
	"time"

	"github.com/laurent22/toml-go/toml"

	"supportlocal/TEDxMileHigh/commands"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"supportlocal/TEDxMileHigh/redis"
)

func init() { commands.Register(command{"manager"}) }

type command struct{ name string }

func (cmd command) Name() string           { return cmd.name }
func (cmd command) CanCreatePidFile() bool { return true }

// periodically cycles through messages
func (cmd command) Run(config toml.Document) {
	debug := config.GetBool("debug") || config.GetBool("manager.debug")

	duration, err := time.ParseDuration(config.GetString("manager.ticker-duration"))
	fatal.If(err)
	ticker := time.NewTicker(duration)

	messageRepo := redis.MessageRepo()

	for _ = range ticker.C {
		message, err := messageRepo.Cycle()
		fatal.If(err)

		if debug {
			log.Printf("manager: cycled to message %d", message.Id)
		}
	}
}
