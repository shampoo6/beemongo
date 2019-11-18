package user_service

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/shampoo6/beemongo/domains"
	"github.com/shampoo6/beemongo/models"
	"github.com/shampoo6/beemongo/models/dto"
	"github.com/shampoo6/beemongo/mongo/cmd/common"
	"github.com/shampoo6/beemongo/mongo/cmd/save"
	"github.com/shampoo6/beemongo/mongo/connection/pool"
	"github.com/shampoo6/beemongo/utils/copy_field"
)

func Insert(dto *dto.UserDto) *domains.User {
	user := new(domains.User)
	copy_field.Copy(dto, user)
	user.Id = ""
	user.Age = int16(dto.Age)
	save.Save(user)
	return user
}

func Update(dto *dto.UserDto) *domains.User {
	user := new(domains.User)
	copy_field.Copy(dto, user)
	user.Age = int16(dto.Age)
	fn := func(db *mgo.Database) interface{} {
		c := db.C("User")
		change := mgo.Change{
			Update:    common.GetUpdateM(user),
			ReturnNew: true,
		}
		// findAndModify操作
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
