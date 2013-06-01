package redis

import (
	"supportlocal/TEDxMileHigh/models"
	"testing"
)

func Test_messageRepo(t *testing.T) {
	var (
		count int

		head  models.Message
		tail  models.Message
		cycle models.Message

		repo = MessageRepo()
		msg1 = models.Message{Author: "bart", Email: "bart@simpsons.com", Comment: "Cowabunga!"}
		msg2 = models.Message{Author: "lisa", Email: "lisa@simpsons.com", Comment: "Bart!"}
		msg3 = models.Message{Author: "apu", Email: "apu@simpsons.com", Comment: "I am so sorry, sir. Please accept five pounds of frozen shrimp."}
	)

	assertSame := func(expected, observed models.Message, comment string) {
		if expected.Id != observed.Id || expected.Author != observed.Author || expected.Email != observed.Email || expected.Comment != observed.Comment {
			t.Fatalf("%s\nexpected: %#v\nobserved: %#v", comment, expected, observed)
		}
	}

	if count, _ = repo.Count(); count != 0 {
		t.Fatalf("wrong count %d", count)
	}

	// msg1 ---

	_ = repo.Save(&msg1)
	if count, _ = repo.Count(); count != 1 {
		t.Fatalf("wrong count %d", count)
	}

	head, _ = repo.Head()
	assertSame(msg1, head, "after msg1, head should be msg1")

	tail, _ = repo.Tail()
	assertSame(msg1, tail, "after msg1, tail should be msg1")

	// msg2 ---

	_ = repo.Save(&msg2)
	if count, _ = repo.Count(); count != 2 {
		t.Fatalf("wrong count %d", count)
	}

	head, _ = repo.Head()
	assertSame(msg1, head, "after msg2, head should be msg1")

	tail, _ = repo.Tail()
	assertSame(msg2, tail, "after msg2, tail should be msg2")

	// msg3 ---

	_ = repo.Save(&msg3)
	if count, _ = repo.Count(); count != 3 {
		t.Fatalf("wrong count %d", count)
	}

	head, _ = repo.Head()
	assertSame(msg1, head, "after msg3, head should be msg1")

	tail, _ = repo.Tail()
	assertSame(msg3, tail, "after msg3, tail should be msg3")

	// ------------- Cycle

	cycle, _ = repo.Cycle()
	assertSame(msg1, cycle, "first cycle should be msg1")

	cycle, _ = repo.Cycle()
	assertSame(msg2, cycle, "second cycle should be msg2")

	cycle, _ = repo.Cycle()
	assertSame(msg3, cycle, "third cycle should be msg3")

	cycle, _ = repo.Cycle()
	assertSame(msg1, cycle, "fourth cycle should be msg1")
}
