package mongo

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
	"time"
)

type InboundMessage struct {
	Id      bson.ObjectId     `bson:"_id" json:"id"`
	Comment string            `bson:"c"   json:"c"`
	Email   string            `bson:"e"   json:"e"`
	Name    string            `bson:"n"   json:"n"`
	Ban     bool              `bson:"ban" json:"ban"`
	Created time.Time         `bson:"cat" json:"cat"`
	Updated time.Time         `bson:"uat" json:"uat"`
	Errors  map[string]string `bson:"-"   json:"-"`
}

func (m *InboundMessage) Valid() bool {
	m.Errors = make(map[string]string)

	if name := strings.TrimSpace(m.Name); len(name) == 0 {
		m.Errors["name"] = "is required"
	}

	if comment := strings.TrimSpace(m.Comment); len(comment) == 0 {
		m.Errors["comment"] = "is required"
	}

	return len(m.Errors) == 0
}

func (m InboundMessage) ToCurrentMessage() CurrentMessage {
	return CurrentMessage{
		Id:      m.Id,
		Author:  m.Name,
		Comment: m.Comment,
	}
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

func (r inboundMessageRepo) Next(id bson.ObjectId) (inboundMessage InboundMessage, err error) {
	key := M{"_id": M{"$gt": id}, "ban": false}

	err = r.collection.Find(key).Sort("_id").Limit(1).One(&inboundMessage)

	return
}

func (r inboundMessageRepo) Save(inboundMessage *InboundMessage) error {
	if !inboundMessage.Valid() {
		return fmt.Errorf("inboundMessage is invalid %#v", inboundMessage.Errors)
	}

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
		iter = r.collection.Find(key).Sort("-_id").Tail(-1)
	)

	for iter.Next(&msg) {
		callback(msg)
	}

	return iter.Close()
}

func InboundMessageRepo() inboundMessageRepo {
	collection := Database.C("inbound_messages")

	return inboundMessageRepo{collection}
}
