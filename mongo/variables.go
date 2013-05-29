package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	Database *mgo.Database

	currentMessageCollection *mgo.Collection

	newObjectId = bson.NewObjectId
	emptyStruct = &struct{}{}
)
