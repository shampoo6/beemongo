package connection

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var _connection *Connection

type Connection struct {
	url    string
	dbName string
	db     *mongo.Database
}

func GetDB() *mongo.Database {
	if _connection != nil {
		return _connection.db
	}
	_connection = new(Connection)
	_connection.url = beego.AppConfig.String("mongodb::url")
	_connection.dbName = beego.AppConfig.String("mongodb::db")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(_connection.url).SetMaxPoolSize(50)) // 连接池
	if err != nil {
		logs.Error("数据库连接失败")
		panic(err)
	}
	_connection.db = client.Database(_connection.dbName)
	return _connection.db
}
