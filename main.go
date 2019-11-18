package main

import (
	"github.com/astaxie/beego"
	_ "github.com/shampoo6/beemongo/conf"
	_ "github.com/shampoo6/beemongo/routers"
)

func main() {

	beego.Run()

	//user := new(domains.User)
	//user.Id = bson.ObjectIdHex("5dcd21e362bb31ac08ecd6ce")
	//user.Name = "Amy66"
	//user.Sex = "Female"
	//user.Age = 18
	//save.Save(user)
	//logs.Debug(user)

	//fn := func(db *mgo.Database) interface{} {
	//	c := db.C("User")
	//	var result []domains.User
	//	c.Find(bson.M{}).All(&result)
	//	return result
	//}
	//
	//result := pool.GetConnectionPool().ExecDbFn(fn).([]domains.User)
	//logs.Debug(result)

	//recoverFunc := beego.BConfig.RecoverFunc
	//beego.BConfig.RecoverFunc = func(c *context.Context){
	//	logs.Debug("呵呵")
	//}
	//myFn := func(c *context.Context) {
	//	logs.Debug("call recoverFunc")
	//	logs.Debug(recoverFunc)
	//	recoverFunc(c)
	//	beego.BConfig.RecoverFunc(c)
	//	logs.Debug("do myself")
	//}
	//logs.Debug(beego.BConfig.RecoverFunc)
	//myFn(nil)
}
