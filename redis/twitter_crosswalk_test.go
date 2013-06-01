package redis

import (
	"testing"
)

func Test_twitterCrosswalk(t *testing.T) {
	var crosswalk = TwitterCrosswalk(MessageRepo())

	id1, _ := crosswalk.MessageIdFor(1)
	id2, _ := crosswalk.MessageIdFor(1)
	id3, _ := crosswalk.MessageIdFor(2)

	if id1 == 0 || id2 == 0 || id3 == 0 {
		t.Fatalf("missing id: %d %d %d", id1, id2, id3)
	}

	if id1 != id2 {
		t.Fatalf("%d != %d", id1, id2)
	}

	if id1 == id3 {
		t.Fatalf("%d == %d", id1, id3)
	}
}
