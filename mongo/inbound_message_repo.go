package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type InboundMessage struct {
	Id      bson.ObjectId `bson:"_id" json:"id"`
	Ban     bool          `bson:"ban" json:"ban"`
	Created time.Time     `bson:"cat" json:"cat"`
	Updated time.Time     `bson:"uat" json:"uat"`
}

type inboundMessageRepo struct {
	collection *mgo.Collection
}

func (r inboundMessageRepo) Ban(id bson.ObjectId) (err error) {

	chg := mgo.Change{
		Update: M{
			"$set": M{
				"ban": true,
				"u":   time.Now(),
			}}}

	_, err = r.collection.FindId(id).Apply(chg, emptyStruct)

	return
}

func (r inboundMessageRepo) Save(inboundMessage *InboundMessage) error {
	// TODO if !inboundMessage.Valid() { return an error }

	if !inboundMessage.Id.Valid() {
		inboundMessage.Id = newObjectId()
	}

	if inboundMessage.Created.IsZero() {
		inboundMessage.Created = time.Now()
	}

	inboundMessage.Updated = time.Now()

	_, err := r.collection.UpsertId(inboundMessage.Id, inboundMessage)

	return err
}

func (r inboundMessageRepo) Tail(callback func(InboundMessage)) error {
	var (
		msg  InboundMessage
		key  = M{"_id": M{"$gt": newObjectId()}}
		iter = r.collection.Find(key).Sort("$natural").Tail(-1)
	)

	for iter.Next(&msg) {
		callback(msg)
	}

	return iter.Close()
}

func InboundMessageRepo(db *mgo.Database) inboundMessageRepo {
	collection := db.C("current_message")

	return inboundMessageRepo{collection}
}
