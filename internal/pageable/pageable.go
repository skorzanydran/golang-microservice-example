package pageable

import (
	"gorm.io/gorm"
	"math"
)

type Page struct {
	Size          int
	Page          int
	Sort          string
	TotalPages    int
	TotalElements int64
	Content       []interface{}
}

type Pageable struct {
	Size int
	Page int
	Sort string
}

func InitPageable(Page int, Size int, Sort string) Pageable {
	return Pageable{Page: Page, Sort: Sort, Size: Size}
}

func (p *Pageable) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pageable) GetLimit() int {
	if p.Size == 0 {
		p.Size = 10
	}
	return p.Size
}

func (p *Pageable) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pageable) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Created desc"
	}
	return p.Sort
}

func Paginate(value interface{}, pageable Pageable, page *Page, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	page.TotalElements = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(page.Size)))
	page.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pageable.GetOffset()).Limit(pageable.GetLimit()).Order(pageable.GetSort())
	}
}
