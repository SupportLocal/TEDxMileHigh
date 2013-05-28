package streamer

import (
	"supportlocal/TEDxMileHigh/commands"
)

func init() { commands.Register(command{"streamer"}) }

type command struct{ name string }

func (cmd command) Name() string        { return cmd.name }
func (cmd command) CreatePidFile() bool { return true }

func (cmd command) Run(args []string) {

	<-make(chan bool) // don't exit
}
