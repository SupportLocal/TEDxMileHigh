package repos

import (
	"supportlocal/TEDxMileHigh/models"
)

type MessageRepo interface {
	NextId() (int, error)

	Count() (int, error)

	Cycle() (models.Message, error)
	Head() (models.Message, error)
	Tail() (models.Message, error)

	Save(msg *models.Message) (err error)
}
