package mongo

import (
	"labix.org/v2/mgo/bson"
)

var (
	newObjectId = bson.NewObjectId

	emptyStruct = &struct{}{}
)
