package models

import "math"

type Pagination struct {
	Page      int
	PrevPage  int
	NextPage  int
	Limit     int
	Offset    int
	TotalPage int
	TotalRow  int64
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 || p.Page < 1 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetPreviosPage() int {
	if p.GetPage() >= p.GetTotalPage() {
		p.PrevPage = p.GetTotalPage() - 1
	}else{
		p.PrevPage = p.GetPage() - 1
	}
	return p.PrevPage
}
func (p *Pagination) GetNextPage() int {
	if p.GetPage() >= p.GetTotalPage() {
		p.NextPage = 0
	}else{
		p.NextPage = p.GetPage() + 1
	}
	return p.NextPage
}

func (p *Pagination) GetTotalPage() int {
	if p.Limit >= int(p.TotalRow) {
		p.TotalPage = 1
	} else {
		p.TotalPage = int(math.Ceil(float64(p.TotalRow) / float64(p.Limit)))
	}
	return p.TotalPage
}
