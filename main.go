package main

import (
	"os"

	"supportlocal/TEDxMileHigh/commands"

	_ "supportlocal/TEDxMileHigh/commands/manager"
	_ "supportlocal/TEDxMileHigh/commands/streamer"
	_ "supportlocal/TEDxMileHigh/commands/usage"
	_ "supportlocal/TEDxMileHigh/commands/website"
)

func main() {
	args := os.Args

	commandName := "usage"

	if len(args) > 1 {
		commandName = args[1]
	}

	commands.Find(commandName).Run(args)
}
