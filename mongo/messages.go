package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Message struct {
	Id      bson.ObjectId `bson:"_id"     json:"id"`
	Author  string        `bson:"author"  json:"author"`
	Comment string        `bson:"comment" json:"comment"`
}

type messagesRepo struct {
	collection *mgo.Collection
}

func (m messagesRepo) Tail(callback func(Message)) error {
	var (
		message Message

		key  = bson.M{"_id": bson.M{"$gt": bson.NewObjectId()}}
		iter = m.collection.Find(key).Sort("$natural").Tail(-1)
	)

	for iter.Next(&message) {
		callback(message)
	}

	return iter.Close()
}

func NewMessagesRepo(db *mgo.Database) messagesRepo {
	collection := db.C("messages")

	err := collection.Create(&mgo.CollectionInfo{
		Capped:   true,
		MaxBytes: 100000,
	})

	if err != nil && err.Error() != "collection already exists" {
		panic(err)
	}

	return messagesRepo{collection}
}
