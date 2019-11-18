package common

import (
	"github.com/globalsign/mgo/bson"
	"reflect"
	"time"
)

// 获取更新操作的bson，其中不会更新id和createTime，自动更新updateTime
func GetUpdateM(iPtr interface{}) bson.M {
	result := bson.M{}
	tElem := reflect.TypeOf(iPtr).Elem()
	vElem := reflect.ValueOf(iPtr).Elem()
	for i := 0; i < tElem.NumField(); i++ {
		field := tElem.Field(i)
		if field.Name == "Id" || field.Name == "CreateTime" {
			continue
		}
		if field.Name == "UpdateTime" {
			result[field.Name] = time.Now().Unix()
			continue
		}
		result[field.Name] = vElem.FieldByName(field.Name).Interface()
	}
	return bson.M{"$set": result}
}
