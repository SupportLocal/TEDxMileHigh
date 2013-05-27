package usage

import (
	"supportlocal/TEDxMileHigh/commands"
)

func init() { commands.Register(command{"usage"}) }

type command struct{ name string }

func (cmd command) Name() string { return cmd.name }

func (cmd command) Run(args []string) {
}
