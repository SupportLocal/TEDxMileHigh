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

	duration, err := time.ParseDuration(config.GetString("manager.cycle-every"))
	fatal.If(err)
	ticker := time.NewTicker(duration)

	messageRepo := redis.MessageRepo()

	for _ = range ticker.C {

		message, err := messageRepo.Cycle()
		if err != nil {
			log.Printf("manager: messageRepo.Cycle failed: %s", err)
			continue
		}

		count, err := messageRepo.Count()
		fatal.If(err)
		log.Printf("manager: cycled to message: %d, total: %d", message.Id, count)
	}
}
