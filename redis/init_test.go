package redis

import (
	redigo "github.com/garyburd/redigo/redis"
	"log"
	"time"
)

func init() {
	ConnectionPool = redigo.Pool{
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	c := ConnectionPool.Get()
	defer c.Close()

	if _, err := c.Do("FLUSHALL"); err != nil {
		log.Fatalf("FLUSHALL failed: %q", err)
	}
}
