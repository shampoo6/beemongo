package common

import (
	"github.com/globalsign/mgo/bson"
	"reflect"
)

func GetUpdateM(iPtr interface{}) bson.M {
	result := bson.M{}
	tElem := reflect.TypeOf(iPtr).Elem()
	vElem := reflect.ValueOf(iPtr).Elem()
	for i := 0; i < tElem.NumField(); i++ {
		field := tElem.Field(i)
		if field.Name == "Id" || field.Name == "CreateTime" {
			continue
		}
		result[field.Name] = vElem.FieldByName(field.Name).Interface()
	}
	return bson.M{"$set": result}
}
