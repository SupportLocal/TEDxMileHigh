package mongo

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
)

type CurrentMessage struct {
	Id      bson.ObjectId     `bson:"_id" json:"id"`
	Author  string            `bson:"a"   json:"author"`
	Comment string            `bson:"c"   json:"comment"`
	Errors  map[string]string `bson:"-"   json:"-"`
}

func (m *CurrentMessage) valid() bool {
	m.Errors = make(map[string]string)

	if id := m.Id; !id.Valid() {
		m.Errors["id"] = "is invalid"
	}

	if author := strings.TrimSpace(m.Author); len(author) == 0 {
		m.Errors["author"] = "is required"
	}

	if comment := strings.TrimSpace(m.Comment); len(comment) == 0 {
		m.Errors["comment"] = "is required"
	}

	return len(m.Errors) == 0
}

type currentMessageRepo struct {
	collection *mgo.Collection
}

func (r currentMessageRepo) Last() (currentMessage CurrentMessage, err error) {
	err = r.collection.Find(nil).Sort("-_id").One(&currentMessage)
	return
}

func (r currentMessageRepo) Save(currentMessage *CurrentMessage) error {
	if !currentMessage.valid() {
		return fmt.Errorf("currentMessage is invalid %#v", currentMessage.Errors)
	}

	_, err := r.collection.UpsertId(currentMessage.Id, currentMessage)

	return err
}

func (r currentMessageRepo) Tail(callback func(CurrentMessage)) error {
	var (
		msg  CurrentMessage
		key  = m{"_id": m{"$gt": newObjectId()}}
		iter = r.collection.Find(key).Sort("$natural").Tail(-1)
	)

	for iter.Next(&msg) {
		callback(msg)
	}

	return iter.Close()
}

func CurrentMessageRepo() currentMessageRepo {
	collection := Database.C("current_message")

	err := collection.Create(&mgo.CollectionInfo{
		Capped:   true,
		MaxBytes: 100000000, // ~95MB
	})

	if err != nil && err.Error() != "collection already exists" {
		panic(err)
	}

	return currentMessageRepo{collection}
}
