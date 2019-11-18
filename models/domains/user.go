package domains

import (
	"context"
	"github.com/shampoo6/beemongo/models"
	"github.com/shampoo6/beemongo/models/dto"
	"github.com/shampoo6/beemongo/mongo/connection"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Document
type User struct {
	Id primitive.ObjectID `bson:"_id"` // id
	// @Index
	CreateTime int64 `bson:"CreateTime"` // 创建时间
	// @Index
	UpdateTime int64 `bson:"UpdateTime"` // 更新时间
	// @Index unique
	Name string `bson:"Name"` // 姓名
	// @Index
	Sex string `bson:"Sex"` // 性别
	// @Index
	Age uint8 `bson:"Age"` // 年龄
}

func Page(page *models.Page, dto *dto.UserDto) models.PageResult {
	col := connection.GetDB().Collection("User")
	result := &[]User{}
	query := dto.GetQuery()
	_options := page.Query(col, query)
	cursor, err := col.Find(context.Background(), query, _options)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var one User
		if err = cursor.Decode(&one); err != nil {
			panic(err)
		}
		*result = append(*result, one)
	}
	list := new([]interface{})
	for _, u := range *result {
		*list = append(*list, u)
	}
	return models.PageResult{PageInfo: *page, Data: *list}
}
