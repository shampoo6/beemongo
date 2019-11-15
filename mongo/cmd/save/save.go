package save

import (
	"beemongo/mongo/cmd/common"
	"beemongo/mongo/connection/pool"
	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"reflect"
	"time"
)

func Save(iPtr interface{}) interface{} {
	// 获取表名
	vElem := reflect.ValueOf(iPtr).Elem()
	collectionName := vElem.Type().Name()
	// 插入前操作

	tElem := reflect.TypeOf(iPtr).Elem()
	objectId, isInsert := preInsert(vElem, tElem)

	fn := func(db *mgo.Database) interface{} {
		c := db.C(collectionName)
		if isInsert {
			if err := c.Insert(iPtr); err != nil {
				logs.Error("数据库插入失败")
				panic(err)
			}
		} else {
			if err := c.UpdateId(objectId, common.GetUpdateM(iPtr)); err != nil {
				logs.Error("数据库修改失败")
				panic(err)
			}
		}
		return iPtr
	}
	return pool.GetConnectionPool().ExecDbFn(fn)
}

// 第一个返回的字符串是id号 返回true 说明是insert操作，否则是update
func preInsert(vElem reflect.Value, tElem reflect.Type) (bson.ObjectId, bool) {
	var isInsert bool = false
	var objectId bson.ObjectId = bson.NewObjectId()
	id := vElem.FieldByName("Id").Interface().(bson.ObjectId)
	time := time.Now().Unix()
	if id.Hex() == "" {
		isInsert = true
		vElem.FieldByName("Id").Set(reflect.ValueOf(objectId))
		vElem.FieldByName("CreateTime").SetInt(time)
		vElem.FieldByName("UpdateTime").SetInt(time)
	} else {
		vElem.FieldByName("UpdateTime").SetInt(time)
		objectId = vElem.FieldByName("Id").Interface().(bson.ObjectId)
	}
	return objectId, isInsert
}
