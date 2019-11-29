package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserView struct {
	Id         primitive.ObjectID `bson:"_id"`
	CreateTime int64
	UpdateTime int64
	Name       string
	Sex        string
	Age        uint8
	Email      string
	Mobile     string
	Address    string
}
