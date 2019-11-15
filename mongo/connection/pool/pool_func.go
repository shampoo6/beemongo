package pool

import "github.com/globalsign/mgo"

// 获取数据库对象
func (cp *ConnectionPool) GetDB() *mgo.Database {
	var db *mgo.Database
	db = cp.session.Copy().DB(cp.dbName)
	return db
}

// 关闭数据库对象
func (cp *ConnectionPool) CloseDB(db *mgo.Database) {
	db.Session.Close()
}

// 执行数据库方法
func (cp *ConnectionPool) ExecDbFn(fn func(db *mgo.Database) interface{}) interface{} {
	db := cp.GetDB()
	result := fn(db)
	cp.CloseDB(db)
	return result
}
