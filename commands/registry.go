package commands

import (
	"log"
)

var registry = make(map[string]Command)

type Command interface {
	Name() string
	CanCreatePidFile() bool
	Run(args []string)
}

func Find(name string) Command {
	var ok bool

	cmd, ok := registry[name]

	if !ok {
		cmd = registry["usage"]
	}

	return cmd
}

func Register(cmd Command) {
	key := cmd.Name()

	if _, ok := registry[key]; ok {
		log.Fatalf("Command name is already in use! %q", key)
	}

	registry[key] = cmd
}
