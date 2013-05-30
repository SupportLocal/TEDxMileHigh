package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type TwitterCrosswalk struct {
	ExternalId int64         `bson:"_id"`
	InternalId bson.ObjectId `bson:"iid"`
	Created    time.Time     `bson:"cat"`
}

type twitterCrosswalkRepo struct {
	collection *mgo.Collection
}

func (r twitterCrosswalkRepo) FindOrCreate(id int64) (twitterCrosswalk TwitterCrosswalk, err error) {

	chg := mgo.Change{
		ReturnNew: true,
		Upsert:    true,
		Update: m{
			"$setOnInsert": m{
				"iid": newObjectId(),
				"cat": time.Now(),
			},
		}}

	_, err = r.collection.FindId(id).Apply(chg, &twitterCrosswalk)

	return
}

func TwitterCrosswalkRepo() twitterCrosswalkRepo {
	collection := Database.C("twitter_crosswalk")

	return twitterCrosswalkRepo{collection}
}
