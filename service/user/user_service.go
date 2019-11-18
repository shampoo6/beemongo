package user_service

import (
	"github.com/shampoo6/beemongo/models"
	"github.com/shampoo6/beemongo/models/domains"
	"github.com/shampoo6/beemongo/models/dto"
	"github.com/shampoo6/beemongo/mongo/cmd/common"
	"github.com/shampoo6/beemongo/utils/copy_field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func DeleteAll(ids []string) int64 {
	return common.DeleteAll("User", ids)
}

func Page(page *models.Page, dto *dto.UserDto) interface{} {
	return domains.Page(page, dto)
}
