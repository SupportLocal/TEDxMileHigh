package commands

import (
	"log"
)

var registry = make(map[string]Command)

type Command interface {
	Name() string
	Run(args []string)
}

func Find(name string) Command {
	cmd, ok := registry[name]

	if !ok {
		log.Fatalf("Command not found! %q", name)
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
