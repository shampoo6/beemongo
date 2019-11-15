package domains

import "github.com/globalsign/mgo/bson"

// @Document
type User struct {
	Id bson.ObjectId `bson:"_id"` // id
	// @Index
	CreateTime int64 `bson:"CreateTime"` // 创建时间
	// @Index
	UpdateTime int64 `bson:"UpdateTime"` // 更新时间
	// @Index unique
	Name string `bson:"Name"` // 姓名
	// @Index
	Sex string `bson:"Sex"` // 性别
	// @Index
	Age int16 `bson:"Age"` // 年龄
}
