package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	redigo "github.com/garyburd/redigo/redis"
	"github.com/laurent22/toml-go/toml"

	"supportlocal/TEDxMileHigh/commands"
	"supportlocal/TEDxMileHigh/redis"

	_ "supportlocal/TEDxMileHigh/commands/manager"
	_ "supportlocal/TEDxMileHigh/commands/streamer"
	_ "supportlocal/TEDxMileHigh/commands/usage"
	_ "supportlocal/TEDxMileHigh/commands/website"
)

const defaultConfigFile = "./etc/TEDxMileHigh.toml"

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", defaultConfigFile, "config file")
	flag.StringVar(&configFile, "c", defaultConfigFile, "config file (shorthand)")
	flag.Parse()

	var (
		parser      toml.Parser
		config      = parser.ParseFile(configFile) // note: ParseFile panics
		commandName = "usage"
	)

	if len(flag.Args()) > 0 {
		commandName = flag.Args()[0]
	}

	command := commands.Find(commandName)

	if command.CanCreatePidFile() {
		fileName := filepath.Join(
			config.GetString("pids"),
			fmt.Sprintf("%s.pid", command.Name()))

		if file, err := os.Create(fileName); err == nil {
			file.WriteString(fmt.Sprintf("%v", os.Getpid()))
			file.Close()
		}
	}

	redis.ConnectionPool = redigo.Pool{
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial("tcp", ":6379")
		},
	}

	command.Run(config)
}
