package redis

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"strconv"
	"supportlocal/TEDxMileHigh/domain/models"
	"supportlocal/TEDxMileHigh/domain/pubsub"
	"supportlocal/TEDxMileHigh/domain/repos"
	_pager "supportlocal/TEDxMileHigh/lib/pager"
)

func MessageRepo() repos.MessageRepo {
	return messageRepo{}
}

type messageRepo struct{}

func (r messageRepo) Blocked() (int, error) {
	return r.llen(blockedListKey)
}

func (r messageRepo) Count() (int, error) {
	return r.llen(activeListKey)
}

func (r messageRepo) PaginateBlocked(pager _pager.Pager) (models.Messages, error) {
	return r.paginate(pager, blockedListKey)
}

func (r messageRepo) PaginatePending(pager _pager.Pager) (models.Messages, error) {
	return r.paginate(pager, activeListKey)
}

func (r messageRepo) Subscribe(channels ...pubsub.Channel) pubsub.Subscription {
	return messageRepoSubscription{
		channels:    channels,
		connection:  redigo.PubSubConn{Conn: ConnectionPool.Get()},
		messageRepo: r,
	}
}

func (r messageRepo) Cycle() (msg models.Message, err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	var id int

	if id, err = redigo.Int(c.Do("RPOPLPUSH", activeListKey, activeListKey)); err != nil {
		return
	}

	if msg, err = r.findById(c, id); err != nil {
		return
	}

	_, err = redigo.Int(c.Do("PUBLISH", pubsub.MessageCycled, msg.Id))

	return
}

func (r messageRepo) Head() (msg models.Message, err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	var values []interface{}

	if values, err = redigo.Values(c.Do("LRANGE", activeListKey, -1, -1)); err != nil {
		return
	}

	return r.findByIdVals(c, values)
}

func (r messageRepo) Tail() (msg models.Message, err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	var values []interface{}

	if values, err = redigo.Values(c.Do("LRANGE", activeListKey, 0, 0)); err != nil {
		return
	}

	return r.findByIdVals(c, values)
}

func (r messageRepo) NextId() (int, error) {
	c := ConnectionPool.Get()
	defer c.Close()

	return redigo.Int(c.Do("INCR", messageIdKey))
}

func (r messageRepo) Block(id int) (err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	c.Send("MULTI")
	c.Send("LPUSH", blockedListKey, id)
	c.Send("LREM", activeListKey, 0, id)
	c.Send("PUBLISH", pubsub.MessageBlocked, id)
	_, err = c.Do("EXEC")

	return
}

func (r messageRepo) Find(id int) (models.Message, error) {
	c := ConnectionPool.Get()
	defer c.Close()
	return r.findById(c, id)
}

func (r messageRepo) Save(msg *models.Message) (err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	isNew := msg.Id == 0

	if isNew {
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
		"e", msg.Email,
	)

	if !isNew {
		c.Send("LREM", activeListKey, 0, msg.Id)
	}

	c.Send("LPUSH", activeListKey, msg.Id)

	if isNew {
		c.Send("PUBLISH", pubsub.MessageCreated, msg.Id)
	} else {
		c.Send("PUBLISH", pubsub.MessageUpdated, msg.Id)
	}

	c.Send("PUBLISH", pubsub.MessageSaved, msg.Id)

	_, err = c.Do("EXEC")

	return
}

func (r messageRepo) allByIdVals(c redigo.Conn, idvals []interface{}) (msgs models.Messages, err error) {
	msgs = make(models.Messages, len(idvals))

	for i, idval := range idvals {
		var msgId int

		if msgId, err = strconv.Atoi(fmt.Sprintf("%s", idval)); err != nil {
			return
		}

		if msgs[i], err = r.findById(c, msgId); err != nil {
			return
		}
	}

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

func (r messageRepo) llen(list string) (int, error) {
	c := ConnectionPool.Get()
	defer c.Close()
	return redigo.Int(c.Do("LLEN", list))
}

func (r messageRepo) paginate(pager _pager.Pager, list string) (messages models.Messages, err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	var (
		totalEntries int

		values []interface{}

		// reverse indexes
		stop  = (pager.Offset() * -1) - 1
		start = stop - pager.PerPage() + 1
	)

	if totalEntries, err = redigo.Int(c.Do("LLEN", list)); err != nil {
		return
	}

	pager.SetTotalEntries(totalEntries)

	if values, err = redigo.Values(c.Do("LRANGE", list, start, stop)); err != nil {
		return
	}

	// reverse values
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}

	return r.allByIdVals(c, values)
}
