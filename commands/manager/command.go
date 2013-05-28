package manager

import (
	"fmt"
	"time"

	"labix.org/v2/mgo"

	"supportlocal/TEDxMileHigh/commands"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"supportlocal/TEDxMileHigh/mongo"
)

func init() { commands.Register(command{"manager"}) }

type command struct{ name string }

func (cmd command) Name() string           { return cmd.name }
func (cmd command) CanCreatePidFile() bool { return true }

// periodically copies inbound messages in to the current message collection
func (cmd command) Run(args []string) {

	session, err := mgo.Dial("localhost")
	fatal.If(err)
	mongo.Database = session.DB("tedx")

	//inboundMessageRepo := mongo.InboundMessageRepo()
	//currentMessageRepo := mongo.CurrentMessageRepo()

	ticker := time.NewTicker(20 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	<-make(chan bool) // don't exit
}
