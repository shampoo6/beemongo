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
	"strings"
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

func Page(page *models.Page, dto *models.UserDto) interface{} {
	fn := func(db *mgo.Database) interface{} {
		c := db.C("User")
		query := bson.M{}
		if strings.Trim(dto.Name, " ") != "" {
			query["Name"] = bson.M{"$regex": "^(\\s|\\S)*" + dto.Name + "(\\s|\\S)*$"}
		}
		if dto.Age > 0 {
			query["Age"] = dto.Age
		}
		if strings.Trim(dto.Sex, " ") != "" {
			query["Sex"] = dto.Sex
		}
		var list []domains.User
		q, total := page.Query(c, query)
		_ = q.All(&list)
		page.SetTotalElement(total)
		var iList []interface{}
		for _, user := range list {
			iList = append(iList, user)
		}
		return models.PageResult{PageInfo: page, Data: &iList}
	}
	return pool.GetConnectionPool().ExecDbFn(fn)
}
