package main

import (
	_ "beemongo/hooks/app/start"
	_ "beemongo/routers"
	"github.com/astaxie/beego"
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
}
