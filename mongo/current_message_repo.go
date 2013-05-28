package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type CurrentMessage struct {
	Id      bson.ObjectId `bson:"_id"     json:"id"`
	Author  string        `bson:"author"  json:"author"`
	Comment string        `bson:"comment" json:"comment"`
}

type currentMessageRepo struct {
	collection *mgo.Collection
}

func (r currentMessageRepo) Tail(callback func(CurrentMessage)) error {
	var (
		msg  CurrentMessage
		key  = M{"_id": M{"$gt": newObjectId()}}
		iter = r.collection.Find(key).Sort("$natural").Tail(-1)
	)

	for iter.Next(&msg) {
		callback(msg)
	}

	return iter.Close()
}

func CurrentMessageRepo(db *mgo.Database) currentMessageRepo {
	collection := db.C("current_message")

	err := collection.Create(&mgo.CollectionInfo{
		Capped:   true,
		MaxBytes: 100000,
	})

	if err != nil && err.Error() != "collection already exists" {
		panic(err)
	}

	return currentMessageRepo{collection}
}
