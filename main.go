package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/laurent22/toml-go/toml"

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

	// load a configuration file
	configFile, _ := filepath.Abs("./TEDxMileHigh.toml") // TODO accept this as a command line arg

	var parser toml.Parser
	config := parser.ParseFile(configFile) // note: ParseFile panics

	command := commands.Find(commandName)

	// TODO if user wants pid-files and the command can create pidfiles
	if command.CanCreatePidFile() {

		fileName := filepath.Join(
			config.GetString("pids"),
			fmt.Sprintf("TEDxMileHigh-%s.pid", command.Name()))

		if file, err := os.Create(fileName); err == nil {
			file.WriteString(fmt.Sprintf("%v", os.Getpid()))
			file.Close()
		}
	}

	command.Run(args, config)
}
