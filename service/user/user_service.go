package user_service

import (
	"beemongo/domains"
	"beemongo/models"
	"beemongo/models/dto"
	"beemongo/mongo/cmd/common"
	"beemongo/mongo/cmd/save"
	"beemongo/mongo/connection/pool"
	"beemongo/utils/copy_field"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

func Insert(dto *dto.UserDto) *domains.User {
	user := new(domains.User)
	copy_field.Copy(dto, user)
	user.Id = bson.ObjectId("")
	user.Age = int16(dto.Age)
	save.Save(user)
	return user
}

func Update(dto *dto.UserDto) *domains.User {
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

func Page(page *models.Page, dto *dto.UserDto) interface{} {
	fn := func(db *mgo.Database) interface{} {
		c := db.C("User")
		//var list []domains.User
		list := make([]interface{}, 0) // 如果不操作返回结果，可以不设置查询返回类型
		q, total := page.Query(c, dto.GetQuery())
		_ = q.All(&list)
		page.SetTotalElement(total)
		//iList := make([]interface{}, 0)
		//for _, user := range list {
		//	iList = append(iList, user)
		//}
		return models.PageResult{PageInfo: *page, Data: list}
	}
	return pool.GetConnectionPool().ExecDbFn(fn)
}
