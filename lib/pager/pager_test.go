package pager

import (
	"net/url"
	"testing"
)

func Test_Parse(t *testing.T) {
	var (
		p Pager
		v url.Values
	)

	// --------------------

	v = url.Values{}
	v.Set("pg", "10")
	v.Set("pp", "100")
	p = Parse(v)

	if p.Page() != 10 {
		t.Fatal(p.Page())
	}

	if p.PerPage() != 100 {
		t.Fatal(p.PerPage())
	}

	if p.TotalEntries() != 0 {
		t.Fatal(p.TotalEntries())
	}

	if p.TotalPages() != 1 {
		t.Fatal(p.TotalPages())
	}

	p.SetTotalEntries(10000)

	if p.TotalEntries() != 10000 {
		t.Fatal(p.TotalEntries())
	}

	if p.TotalPages() != 100 {
		t.Fatal(p.TotalPages())
	}

	if p.Prev() != 9 {
		t.Fatal(p.Prev())
	}

	if p.Next() != 11 {
		t.Fatal(p.Next())
	}

	// --------------------

	v = url.Values{}
	v.Set("pg", "1")
	v.Set("pp", "10")
	p = Parse(v)
	p.SetTotalEntries(11)

	if p.HasPrev() {
		t.Fatal(p.HasPrev())
	}

	if p.Next() != 2 {
		t.Fatal(p.Next())
	}

	if !p.HasNext() {
		t.Fatal(p.HasNext())
	}

	if p.TotalPages() != 2 {
		t.Fatal(p.TotalPages())
	}

	// --------------------

	v = url.Values{}
	v.Set("pg", "2")
	v.Set("pp", "10")
	p = Parse(v)
	p.SetTotalEntries(11)

	if !p.HasPrev() {
		t.Fatal(p.HasPrev())
	}

	if p.HasNext() {
		t.Fatal(p.HasNext())
	}

}
