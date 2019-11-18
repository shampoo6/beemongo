package user_service

import (
	"context"
	"github.com/shampoo6/beemongo/domains"
	"github.com/shampoo6/beemongo/models"
	"github.com/shampoo6/beemongo/models/dto"
	"github.com/shampoo6/beemongo/mongo/cmd/common"
	"github.com/shampoo6/beemongo/mongo/connection"
	"github.com/shampoo6/beemongo/utils/copy_field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// todo 数据库方法应该移动到model中

func Insert(dto *dto.UserDto) *domains.User {
	user := new(domains.User)
	copy_field.Copy(dto, user)
	user.Id = primitive.ObjectID{} // 清空前端传来的id值
	user.Age = uint8(dto.Age)
	common.Save(user)
	return user
}

func Update(dto *dto.UserDto) *domains.User {
	user := new(domains.User)
	copy_field.Copy(dto, user)
	user.Age = uint8(dto.Age)
	common.SimpleFindAndModify(user, user)
	return user
}

func Page(page *models.Page, dto *dto.UserDto) interface{} {
	col := connection.GetDB().Collection("User")
	result := &[]domains.User{}
	query := dto.GetQuery()
	_options := page.Query(col, query)
	cursor, err := col.Find(context.Background(), query, _options)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var one domains.User
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
