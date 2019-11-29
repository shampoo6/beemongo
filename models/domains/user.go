package domains

import (
	"context"
	"github.com/shampoo6/beemongo/models"
	"github.com/shampoo6/beemongo/models/dto"
	"github.com/shampoo6/beemongo/mongo/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @document
type User struct {
	Id primitive.ObjectID `bson:"_id"` // id
	// @index
	CreateTime int64 `bson:"CreateTime"` // 创建时间
	// @index
	UpdateTime int64 `bson:"UpdateTime"` // 更新时间
	// @index unique
	Name string `bson:"Name"` // 姓名
	// @index
	Sex string `bson:"Sex"` // 性别
	// @index
	Age uint8 `bson:"Age"` // 年龄
}

func Page(page *models.Page, _dto *dto.UserDto) models.PageResult {
	col := connection.GetDB().Collection("UserView")
	result := &[]dto.UserView{}
	query := _dto.GetQuery()
	_options := page.Query(col, query)
	cursor, err := col.Find(context.Background(), query, _options)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var one dto.UserView
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

func Update(ctx context.Context, id string, update bson.M) {
	collection := connection.GetDB().Collection("User")
	objectId, _ := primitive.ObjectIDFromHex(id)
	update = bson.M{"$set": update}
	_, e := collection.UpdateOne(ctx, bson.M{"_id": objectId}, update)
	if e != nil {
		panic(e)
	}
}
