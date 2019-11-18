package start

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/shampoo6/beemongo/mongo/connection/pool"
	"github.com/shampoo6/beemongo/mongo/scanner"
	"path/filepath"
)

// app初始化
func init() {
	beego.AddAPPStartHook(initLog)
	beego.AddAPPStartHook(initMongo)
}

func initLog() error {
	_ = logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	return nil
}

func initMongo() error {
	path, err := filepath.Abs("domains")
	if err != nil {
		logs.Error("自动建表，扫描目录异常")
		panic(err)
	}
	logs.Info("自动建表，扫描目录：")
	logs.Info(path)
	pool.GetConnectionPool()
	scanner.ScanDir(path)
	return nil
}
