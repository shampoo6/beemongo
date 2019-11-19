package domains

import (
	"context"
	"github.com/shampoo6/beemongo/models/dto"
	"github.com/shampoo6/beemongo/mongo/cmd/common"
	"github.com/shampoo6/beemongo/mongo/connection"
	"github.com/shampoo6/beemongo/utils/copy_field"
	"github.com/shampoo6/beemongo/utils/string_util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// @document
type UserInfo struct {
	Id primitive.ObjectID `bson:"_id"`
	// @index unique
	UserId primitive.ObjectID `bson:"UserId"`
	// @index
	Email string `bson:"Email"`
	// @index
	Mobile string `bson:"Mobile"`
	// @index
	Address string `bson:"Address"`
	// @index
	CreateTime int64 `bson:"CreateTime"`
	// @index
	UpdateTime int64 `bson:"UpdateTime"`
}

func Insert(dto *dto.UserInfoDto) interface{} {
	db := connection.GetDB()
	_ = db.Client().UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		e := sessionContext.StartTransaction()
		if e != nil {
			return e
		}
		func() {
			userInfo := new(UserInfo)
			copy_field.Copy(dto, userInfo)
			userInfo.UserId, _ = primitive.ObjectIDFromHex(dto.UserId)
			common.Save(sessionContext, userInfo)
			if string_util.HasText(dto.Name) {
				Update(sessionContext, dto.UserId, bson.M{"Name": dto.Name})
			}
		}()
		if err := recover(); err != nil {
			_ = sessionContext.AbortTransaction(sessionContext)
			panic(err)
		}
		_ = sessionContext.CommitTransaction(sessionContext)
		return nil
	})
	collection := db.Collection("UserView")
	objectId, _ := primitive.ObjectIDFromHex(dto.UserId)
	singleResult := collection.FindOne(context.Background(), bson.M{"_id": objectId})
	var result map[string]interface{}
	_ = singleResult.Decode(&result)
	return result
}
