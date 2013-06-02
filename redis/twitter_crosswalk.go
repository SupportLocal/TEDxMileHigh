package redis

import (
	redigo "github.com/garyburd/redigo/redis"
	"supportlocal/TEDxMileHigh/domain/repos"
)

type twitterCrosswalk struct {
	messageRepo repos.MessageRepo
}

func (cw twitterCrosswalk) MessageIdFor(twitterId int64) (messageId int, err error) {
	c := ConnectionPool.Get()
	defer c.Close()

	var key = tweetKey(twitterId)
	var exists bool

	if exists, err = redigo.Bool(c.Do("HEXISTS", key, "messageId")); err != nil {
		return
	}

	if !exists { // consume a message id, and race to set it
		messageId, err = cw.messageRepo.NextId()
		if _, err = c.Do("HSETNX", key, "messageId", messageId); err != nil {
			return
		}
	}

	// someone has set the messageId by now (may not have been us); look it up
	messageId, err = redigo.Int(c.Do("HGET", key, "messageId"))

	return
}

func TwitterCrosswalk(messageRepo repos.MessageRepo) repos.TwitterCrosswalk {
	return twitterCrosswalk{messageRepo}
}
