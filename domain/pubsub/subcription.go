package pubsub

import (
	"supportlocal/TEDxMileHigh/domain/models"
)

type Subscription interface {
	Receive() (Channel, models.Message, error)
	Unsubscribe() error
}
