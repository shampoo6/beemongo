package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"math"
)

type PageResult struct {
	PageInfo Page
	Data     []interface{}
}

type Page struct {
	Page         int
	Size         int
	TotalPage    int
	TotalElement int
	Sort         []string
}

func (p *Page) initPage() {
	if p.Page < 0 {
		p.Page = 0
	}
	if p.Size <= 0 {
		p.Size = 20
	}
}

func (p *Page) Query(c *mgo.Collection, query bson.M) (*mgo.Query, int) {
	p.initPage()
	find := c.Find(query)
	count, _ := find.Count()
	result := find.Skip(p.Page * p.Size).Limit(p.Size)
	// 例如：result.Sort("UpdateTime") 已UpdateTime字段升序排序
	// result.Sort("-UpdateTime") 已UpdateTime字段降序排序
	if p.Sort != nil {
		for _, sort := range p.Sort {
			result = result.Sort(sort)
		}
	}
	return result, count
}

func (p *Page) SetTotalElement(total int) {
	p.initPage()
	p.TotalElement = total
	p.TotalPage = int(math.Ceil(float64(total) / float64(p.Size)))
}
