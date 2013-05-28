package streamer

import (
	"github.com/laurent22/toml-go/toml"
	"supportlocal/TEDxMileHigh/commands"
)

func init() { commands.Register(command{"streamer"}) }

type command struct{ name string }

func (cmd command) Name() string           { return cmd.name }
func (cmd command) CanCreatePidFile() bool { return true }

func (cmd command) Run(args []string, config toml.Document) {

	<-make(chan bool) // don't exit
}
