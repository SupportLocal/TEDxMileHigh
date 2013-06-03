package repos

import (
	"supportlocal/TEDxMileHigh/domain/models"
	"supportlocal/TEDxMileHigh/domain/pubsub"
	"supportlocal/TEDxMileHigh/lib/pager"
)

type MessageRepo interface {
	NextId() (int, error)

	Count() (int, error)

	Paginate(pager.Pager) (models.Messages, error)

	Subscribe(...pubsub.Channel) pubsub.Subscription

	Cycle() (models.Message, error)
	Head() (models.Message, error)
	Tail() (models.Message, error)

	Find(int) (models.Message, error)
	Save(msg *models.Message) (err error)
}
