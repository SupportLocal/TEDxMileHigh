package pager

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strconv"

	"github.com/levicook/go-detect"
)

const (
	defaultPage    = 1
	minPage        = 1
	defaultPerPage = 30
	maxPerPage     = 500
)

type Pager interface {
	Page() int
	PerPage() int

	TotalEntries() int
	SetTotalEntries(int)

	HasNext() bool
	HasPrev() bool
	Next() int
	Prev() int

	TotalPages() int

	Offset() int
	Skip() int
	Limit() string
}

type pager struct {
	page, perPage, totalEntries int
}

func (p pager) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HasNext      bool `json:"hasNext"`
		HasPrev      bool `json:"hasPrev"`
		Next         int  `json:"next"`
		Page         int  `json:"page"`
		PerPage      int  `json:"perPage"`
		Prev         int  `json:"prev"`
		TotalEntries int  `json:"totalEntries"`
		TotalPages   int  `json:"totalPages"`
	}{
		HasNext:      p.HasNext(),
		HasPrev:      p.HasPrev(),
		Next:         p.Next(),
		Page:         p.Page(),
		PerPage:      p.PerPage(),
		Prev:         p.Prev(),
		TotalEntries: p.TotalEntries(),
		TotalPages:   p.TotalPages(),
	})
}

func (p pager) Page() int {
	return p.page
}

func (p pager) PerPage() int {
	return p.perPage
}

func (p pager) TotalEntries() int {
	return p.totalEntries
}

func (p *pager) SetTotalEntries(t int) {
	if t < 0 {
		t = 0
	}
	p.totalEntries = t
}

func (p pager) HasPrev() bool {
	return p.Prev() > 0
}

func (p pager) Prev() int {
	return p.Page() - 1
}

func (p pager) HasNext() bool {
	return p.Next() <= p.TotalPages()
}

func (p pager) Next() int {
	return p.Page() + 1
}

func (p pager) TotalPages() (totalPages int) {

	if p.totalEntries == 0 {
		totalPages = 1
	} else {
		ceil := math.Ceil(float64(p.totalEntries) / float64(p.perPage))
		totalPages = int(ceil)
	}
	return
}

func (p pager) Skip() int {
	return p.Offset()
}

func (p pager) Offset() int {
	return (p.page - 1) * p.perPage
}

func (p pager) Limit() string {
	return fmt.Sprintf("%v, %v", p.Offset(), p.PerPage())
}

func New(page, perPage int) Pager {
	if perPage > maxPerPage {
		perPage = maxPerPage
	}

	if page < minPage {
		page = defaultPage
	}

	return &pager{
		page:    page,
		perPage: perPage,
	}
}

func Parse(vals url.Values) Pager {

	var perPage, page int
	var parseError error

	perPage, parseError = strconv.Atoi(detect.String(vals.Get("pp"), "30"))
	if parseError != nil {
		perPage = defaultPerPage
	}

	page, parseError = strconv.Atoi(detect.String(vals.Get("pg"), "1"))
	if parseError != nil {
		page = defaultPage
	}

	return New(page, perPage)
}
