package redis

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"strconv"
	"supportlocal/TEDxMileHigh/domain/models"
	"supportlocal/TEDxMileHigh/domain/pubsub"
)

type messageRepoSubscription struct {
	channels    []pubsub.Channel
	connection  redigo.PubSubConn
	messageRepo messageRepo
}

func (s messageRepoSubscription) Receive() (channel pubsub.Channel, message models.Message, err error) {

	for _, cn := range s.channels {
		if err = s.connection.Subscribe(cn); err != nil {
			return
		}
	}

	for {
		switch value := s.connection.Receive().(type) {

		case redigo.Message:
			channel = pubsub.ChannelFor(value.Channel)
			if message.Id, err = strconv.Atoi(fmt.Sprintf("%s", value.Data)); err != nil {
				return
			}
			message, err = s.messageRepo.Find(message.Id)
			return

		case error:
			err = value
			return
		}
	}
}

func (s messageRepoSubscription) Unsubscribe() (err error) {
	return s.connection.Close()
}
