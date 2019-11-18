package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"strings"
)

type PageResult struct {
	PageInfo Page
	Data     []interface{}
}

type Page struct {
	Page         int64
	Size         int64
	TotalPage    int64
	TotalElement int64
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

func (p *Page) Query(c *mongo.Collection, query bson.M) *options.FindOptions {
	p.initPage()
	_options := options.Find().SetSkip(p.Page * p.Size).SetLimit(p.Size)
	if p.Sort != nil {
		sortMap := bson.M{}
		for _, sort := range p.Sort {
			splits := strings.Split(sort, ",")
			if strings.Trim(splits[1], " ") == "asc" {
				sortMap[strings.Trim(splits[0], " ")] = 1
			} else if strings.Trim(splits[1], " ") == "desc" {
				sortMap[strings.Trim(splits[0], " ")] = -1
			}
		}
		_options = _options.SetSort(sortMap)
	}

	total, err := c.CountDocuments(context.Background(), query)
	if err != nil {
		panic(err)
	}
	p.SetTotalElement(total)
	return _options
	//var cursor *mongo.Cursor
	//cursor, err = c.Find(context.Background(), query, _options)
	//if err != nil {
	//	panic(err)
	//}
	//defer cursor.Close(context.Background())
	//for cursor.Next(context.Background()) {
	//	var one interface{}
	//	if err = cursor.Decode(&one); err != nil {
	//		panic(err)
	//	}
	//	*result = append(*result, one)
	//}
	//
	//return PageResult{*p, *result}
}

func (p *Page) SetTotalElement(total int64) {
	p.initPage()
	p.TotalElement = total
	p.TotalPage = int64(math.Ceil(float64(total) / float64(p.Size)))
}
