package pool

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo"
)

var connectionPool *ConnectionPool

type ConnectionPool struct {
	session *mgo.Session
	dbName  string
}

func initSession() *ConnectionPool {
	connectionPool = new(ConnectionPool)
	connectionUrl := beego.AppConfig.String("mongodb::url")
	connectionPool.dbName = beego.AppConfig.String("mongodb::db")
	logs.Info("数据库名：%s", connectionPool.dbName)
	s, err := mgo.Dial(connectionUrl)
	connectionPool.session = s
	if err != nil {
		logs.Error(`数据库连接失败`)
		panic(err)
	}
	connectionPool.session.SetMode(mgo.Monotonic, true)
	logs.Info(`数据库连接成功`)
	return connectionPool
}

func GetConnectionPool() *ConnectionPool {
	if connectionPool == nil {
		return initSession()
	}
	return connectionPool
}
