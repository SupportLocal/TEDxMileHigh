package redis

import (
	"supportlocal/TEDxMileHigh/models"
)

type redisMessage struct {
	Id      int    `redis:"id"`
	Author  string `redis:"a"`
	Comment string `redis:"c"`
	Email   string `redis:"e"`
}

func (rm redisMessage) toMessage() models.Message {
	return models.Message{
		Id:      rm.Id,
		Author:  rm.Author,
		Comment: rm.Comment,
		Email:   rm.Email,
	}
}
