package usage

import (
	"log"
	"supportlocal/TEDxMileHigh/commands"
)

func init() { commands.Register(command{"usage"}) }

type command struct{ name string }

func (cmd command) Name() string { return cmd.name }

func (cmd command) Run(args []string) {
	// TODO list commands from command registry
	log.Println("You're doing it wrong.")
}
