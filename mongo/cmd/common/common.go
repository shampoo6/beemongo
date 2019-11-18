package common

import (
	"context"
	"github.com/shampoo6/beemongo/mongo/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"time"
)

// 将对象直接入库，如果id不存在就insert，存在就update
// iPtr 就是需要保存的结构指针
func Save(iPtr interface{}) interface{} {
	// 获取表名
	vElem := reflect.ValueOf(iPtr).Elem()
	collectionName := vElem.Type().Name()
	// 插入前操作

	objectId, isInsert := preInsert(vElem)

	collection := connection.GetDB().Collection(collectionName)
	var err error
	if isInsert {
		_, err = collection.InsertOne(context.Background(), iPtr)
	} else {
		_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objectId}, GetUpdateM(iPtr))
	}
	if err != nil {
		panic(err)
	}
	return iPtr
}

// iPtr 就是需要更新的结构指针
func FindAndModify(collectionName string, filter interface{}, iPtr interface{}, result interface{}) {
	collection := connection.GetDB().Collection(collectionName)
	after := options.After
	_result := collection.FindOneAndUpdate(context.Background(), filter, GetUpdateM(iPtr), &options.FindOneAndUpdateOptions{ReturnDocument: &after})
	err := _result.Decode(result)
	if err != nil {
		panic(err)
	}
}

// 简易更新，默认更新条件为匹配id，集合为iPtr的类型
func SimpleFindAndModify(iPtr interface{}, result interface{}) {
	// 获取表名
	vElem := reflect.ValueOf(iPtr).Elem()
	collectionName := vElem.Type().Name()
	field := vElem.FieldByName("Id")
	FindAndModify(collectionName, bson.M{"_id": field.Interface()}, iPtr, result)
}

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

// 第一个返回的字符串是id号 返回true 说明是insert操作，否则是update
func preInsert(vElem reflect.Value) (primitive.ObjectID, bool) {
	var isInsert bool = false
	var objectId = primitive.NewObjectID()
	id := vElem.FieldByName("Id").Interface().(primitive.ObjectID)
	_time := time.Now().Unix()
	if id.IsZero() {
		isInsert = true
		vElem.FieldByName("Id").Set(reflect.ValueOf(objectId))
		vElem.FieldByName("CreateTime").SetInt(_time)
		vElem.FieldByName("UpdateTime").SetInt(_time)
	} else {
		vElem.FieldByName("UpdateTime").SetInt(_time)
		objectId = vElem.FieldByName("Id").Interface().(primitive.ObjectID)
	}
	return objectId, isInsert
}
