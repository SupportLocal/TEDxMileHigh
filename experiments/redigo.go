package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"supportlocal/TEDxMileHigh/lib/fatal"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	// This goroutine receives and prints pushed notifications from the server.
	// The goroutine exits when the connection is unsubscribed from all
	// channels or there is an error.
	go func() {
		defer func() {
			// This group never exits.
			log.Println("gr1 waitGroup.Done()")
			waitGroup.Done()
		}()

		con, err := redis.Dial("tcp", ":6379")
		fatal.If(err)
		psc := redis.PubSubConn{con}

		for {
			switch n := psc.Receive().(type) {
			case redis.Message:
				fmt.Printf("Message: %s %s\n", n.Channel, n.Data)
			case redis.PMessage:
				fmt.Printf("PMessage: %s %s %s\n", n.Pattern, n.Channel, n.Data)
			case redis.Subscription:
				fmt.Printf("Subscription: %s %s %d\n", n.Kind, n.Channel, n.Count)
				if n.Count == 0 {
					return
				}
			case error:
				fmt.Printf("Error: %q\n", n)
				return
			}
		}
	}()

	// This goroutine manages subscriptions for the connection.
	go func() {
		defer func() {
			log.Println("gr2 waitGroup.Done()")
			waitGroup.Done()
		}()

		con, err := redis.Dial("tcp", ":6379")
		fatal.If(err)
		psc := redis.PubSubConn{con}

		psc.Subscribe("example")
		psc.PSubscribe("p*")

		// The following function calls publish a message using another
		// connection to the Redis server.
		con.Do("PUBLISH", "example", "hello")
		con.Do("PUBLISH", "example", "world")
		con.Do("PUBLISH", "pexample", "foo")
		con.Do("PUBLISH", "pexample", "bar")

		// Unsubscribe from all connections. This will cause the receiving
		// goroutine to exit.
		psc.Unsubscribe()
		psc.PUnsubscribe()
	}()

	waitGroup.Wait()
}
