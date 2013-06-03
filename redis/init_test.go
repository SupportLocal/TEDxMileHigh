package redis

import (
	redigo "github.com/garyburd/redigo/redis"
	"log"
)

func setupTest() {
	ConnectionPool = redigo.Pool{
		Dial: func() (c redigo.Conn, err error) {
			if c, err = redigo.Dial("tcp", ":6379"); err == nil {
				_, err = c.Do("SELECT", 10)
			}
			return c, err
		},
	}

	c := ConnectionPool.Get()
	defer c.Close()

	if _, err := c.Do("FLUSHDB"); err != nil {
		log.Fatalf("FLUSHDB failed: %q", err)
	}
}

func teardownTest() {
	ConnectionPool.Close()
}
