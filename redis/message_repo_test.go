package redis

import (
	"supportlocal/TEDxMileHigh/domain/models"
	_pager "supportlocal/TEDxMileHigh/lib/pager"
	"testing"
)

func Test_messageRepo(t *testing.T) {
	setupTest()
	defer teardownTest()

	var (
		count int

		head  models.Message
		tail  models.Message
		cycle models.Message

		pager _pager.Pager
		page  models.Messages

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

	// ------------- Paginate

	// page 1,10 ---
	pager = _pager.New(1, 10)
	page, _ = repo.Paginate(pager)

	if pager.TotalEntries() != 3 {
		t.Fatalf("wrong pager.TotalEntries.\nobserved: %d \nexpected: %d", pager.TotalEntries(), 3)
	}

	if len(page) != 3 {
		t.Fatalf("wrong page len: %d", len(page))
	}

	assertSame(msg1, page[0], "page[0] for pager(1,10) should be msg1")
	assertSame(msg2, page[1], "page[1] for pager(1,10) should be msg2")
	assertSame(msg3, page[2], "page[2] for pager(1,10) should be msg3")

	// page 1,1 ---
	pager = _pager.New(1, 1)
	page, _ = repo.Paginate(pager)

	if len(page) != 1 {
		t.Fatalf("wrong page len: %d", len(page))
	}
	assertSame(msg1, page[0], "page[0] for pager(1,1) should be msg1")

	// page 2,1 ---
	pager = _pager.New(2, 1)
	page, _ = repo.Paginate(pager)

	if len(page) != 1 {
		t.Fatalf("wrong page len: %d", len(page))
	}
	assertSame(msg2, page[0], "page[0] for pager(2,1) should be msg2")

	// page 3,1 ---
	pager = _pager.New(3, 1)
	page, _ = repo.Paginate(pager)

	if len(page) != 1 {
		t.Fatalf("wrong page len: %d", len(page))
	}
	assertSame(msg3, page[0], "page[0] for pager(3,1) should be msg3")

	// page 3,2 ---
	/* TODO more pager bugs!!
	pager = _pager.New(3, 2)
	page, _ = repo.Paginate(pager)

	if len(page) != 1 {
		t.Fatalf("wrong page len: %d", len(page))
	}
	assertSame(msg3, page[0], "page[0] for pager(3,2) should be msg3")
	*/

	// page 4,1 ---
	pager = _pager.New(4, 1)
	page, _ = repo.Paginate(pager)

	if len(page) != 0 {
		t.Fatalf("wrong page len: %d", len(page))
	}

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
