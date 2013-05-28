package main

import (
	"fmt"
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

	command := commands.Find(commandName)

	// TODO if user wants pid-files and the command can create pidfiles
	if command.CanCreatePidFile() {
		fileName := fmt.Sprintf("./TEDxMileHigh-%s.pid", command.Name())
		if file, err := os.Create(fileName); err == nil {
			file.WriteString(fmt.Sprintf("%v", os.Getpid()))
			file.Close()
		}
	}

	command.Run(args)
}
