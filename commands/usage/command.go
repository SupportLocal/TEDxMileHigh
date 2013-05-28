package usage

import (
	"github.com/laurent22/toml-go/toml"
	"log"
	"supportlocal/TEDxMileHigh/commands"
)

func init() { commands.Register(command{"usage"}) }

type command struct{ name string }

func (cmd command) Name() string           { return cmd.name }
func (cmd command) CanCreatePidFile() bool { return false }

func (cmd command) Run(args []string, config toml.Document) {

	// TODO list commands from command registry
	log.Println("You're doing it wrong.")
}
