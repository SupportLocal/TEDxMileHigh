package repos

import (
	"supportlocal/TEDxMileHigh/domain/models"
	"supportlocal/TEDxMileHigh/domain/pubsub"
	"supportlocal/TEDxMileHigh/lib/pager"
)

type MessageRepo interface {
	NextId() (int, error)

	Blocked() (int, error) // todo rename: CountBlocked
	Count() (int, error)

	PaginateBlocked(pager.Pager) (models.Messages, error)
	PaginatePending(pager.Pager) (models.Messages, error)

	Subscribe(...pubsub.Channel) pubsub.Subscription

	Cycle() (models.Message, error)
	Head() (models.Message, error)
	Tail() (models.Message, error)

	Block(int) error
	// todo add: Restore(int)

	Find(int) (models.Message, error)
	Save(msg *models.Message) (err error)
}
