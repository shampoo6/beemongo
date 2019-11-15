package user_service

import (
	"beemongo/domains"
	"beemongo/models"
	"beemongo/mongo/cmd/common"
	"beemongo/mongo/cmd/save"
	"beemongo/mongo/connection/pool"
	"beemongo/utils/copy_field"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

func Insert(dto *models.UserDto) *domains.User {
	user := new(domains.User)
	copy_field.Copy(dto, user)
	user.Id = bson.ObjectId("")
	user.Age = int16(dto.Age)
	save.Save(user)
	return user
}

func Update(dto *models.UserDto) *domains.User {
	user := new(domains.User)
	copy_field.Copy(dto, user)
	user.Age = int16(dto.Age)
	user.UpdateTime = time.Now().Unix()
	fn := func(db *mgo.Database) interface{} {
		c := db.C("User")
		change := mgo.Change{
			Update:    common.GetUpdateM(user),
			ReturnNew: true,
		}
		_, _ = c.Find(bson.M{"_id": user.Id}).Apply(change, &user)
		return user
	}
	return pool.GetConnectionPool().ExecDbFn(fn).(*domains.User)
}
