package start

import (
	goContext "context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/shampoo6/beemongo/mongo/connection"
	"github.com/shampoo6/beemongo/mongo/scanner"
	"go.mongodb.org/mongo-driver/bson"
	"path/filepath"
)

// app初始化
func init() {
	beego.AddAPPStartHook(initLog)
	beego.AddAPPStartHook(initMongo)
	beego.AddAPPStartHook(initViews)
	//beego.AddAPPStartHook(filters.InitFilters)
}

func initLog() error {
	_ = logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	return nil
}

func initMongo() error {
	path, err := filepath.Abs("models/domains")
	if err != nil {
		logs.Error("自动建表，扫描目录异常")
		panic(err)
	}
	logs.Info("自动建表，扫描目录：")
	logs.Info(path)
	scanner.ScanDir(path)
	return nil
}

func initViews() error {

	initUserView := func() {
		// 创建视图
		cmd := bson.M{
			"create": "UserView",
			"viewOn": "User",
			"pipeline": bson.A{
				bson.M{
					"$lookup": bson.M{
						"from":         "UserInfo",
						"localField":   "_id",
						"foreignField": "UserId",
						"as":           "my_view",
					},
				},
				bson.M{
					"$replaceRoot": bson.M{
						"newRoot": bson.M{
							"$mergeObjects": bson.A{
								bson.M{
									"$arrayElemAt": bson.A{
										"$my_view", 0,
									},
								},
								"$$ROOT",
							},
						},
					},
				},
				bson.M{
					"$project": bson.M{
						"my_view": 0,
						"UserId":  0,
					},
				},
			},
		}
		connection.GetDB().RunCommand(goContext.Background(), cmd)
		logs.Info("创建UserView")
	}

	initUserView()

	return nil
}
