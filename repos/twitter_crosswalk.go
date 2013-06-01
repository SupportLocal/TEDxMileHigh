package repos

type TwitterCrosswalk interface {
	MessageIdFor(twitterId int64) (int, error)
}
