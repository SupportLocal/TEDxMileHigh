package redis

import (
	"fmt"
)

const (
	messageIdKey   = "message-id"
	activeListKey  = "messages-active"
	blockedListKey = "messages-blocked"
)

func messageKey(id int) string {
	return fmt.Sprintf("message-%v", id)
}

func tweetKey(id int64) string {
	return fmt.Sprintf("tweet-%v", id)
}
