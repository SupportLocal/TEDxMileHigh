package redis

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"strconv"
	"supportlocal/TEDxMileHigh/models"
	"supportlocal/TEDxMileHigh/repos"
)

func MessageRepo() repos.MessageRepo {
	return messageRepo{}
}

type messageRepo struct{}

func (r messageRepo) Count() (int, error) {
	c := ConnectionPool.Get()
	defer c.Close()

	return redigo.Int(c.Do("LLEN", messageListKey))
}

func (r messageRepo) Cycle() (msg models.Message, err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	var id int

	if id, err = redigo.Int(c.Do("RPOPLPUSH", messageListKey, messageListKey)); err != nil {
		return
	}

	return r.findById(c, id)
}

func (r messageRepo) Head() (msg models.Message, err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	var values []interface{}

	if values, err = redigo.Values(c.Do("LRANGE", messageListKey, -1, -1)); err != nil {
		return
	}

	return r.findByIdVals(c, values)
}

func (r messageRepo) Tail() (msg models.Message, err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	var values []interface{}

	if values, err = redigo.Values(c.Do("LRANGE", messageListKey, 0, 0)); err != nil {
		return
	}

	return r.findByIdVals(c, values)
}

func (r messageRepo) NextId() (int, error) {
	c := ConnectionPool.Get()
	defer c.Close()

	return redigo.Int(c.Do("INCR", messageIdKey))
}

func (r messageRepo) Save(msg *models.Message) (err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	if msg.Id == 0 {
		if msg.Id, err = redigo.Int(c.Do("INCR", messageIdKey)); err != nil {
			return
		}
	}

	c.Send("MULTI")

	c.Send(
		"HMSET", messageKey(msg.Id),
		"id", msg.Id,
		"a", msg.Author,
		"c", msg.Comment,
	)

	c.Send("LPUSH", messageListKey, msg.Id)

	_, err = c.Do("EXEC")

	return
}

func (r messageRepo) findById(c redigo.Conn, id int) (msg models.Message, err error) {
	var (
		values          []interface{}
		internalMessage redisMessage
	)

	if values, err = redigo.Values(c.Do("HGETALL", messageKey(id))); err != nil {
		return
	}

	if err = redigo.ScanStruct(values, &internalMessage); err != nil {
		return
	}

	msg = internalMessage.toMessage()

	return
}

func (r messageRepo) findByIdVals(c redigo.Conn, idvals []interface{}) (msg models.Message, err error) {
	for _, idval := range idvals {
		if msg.Id, err = strconv.Atoi(fmt.Sprintf("%s", idval)); err != nil {
			return
		}
	}
	return r.findById(c, msg.Id)
}
