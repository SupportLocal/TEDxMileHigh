package redis

import (
	redigo "github.com/garyburd/redigo/redis"
)

var ConnectionPool redigo.Pool
