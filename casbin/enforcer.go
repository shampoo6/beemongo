package casbin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
	"github.com/casbin/mongodb-adapter"
	"time"
)

var e *casbin.Enforcer

func init() {

	m := model.Model{}
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)")

	url := beego.AppConfig.String("mongodb::url")
	a := mongodbadapter.NewAdapter(url)

	e, _ = casbin.NewEnforcer(m, a)

	// Load the policy from DB.
	err := e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	// Modify the policy.
	_, _ = e.AddPolicy("free", "/casbin/obj/write/:id", "GET")
	_, _ = e.AddPolicy("user", "/casbin/obj/write", "GET|POST")
	_, _ = e.AddRoleForUser("user", "free")
	_, _ = e.AddRoleForUser("5dd273b1629f127417f910ee", "user")
	// e.RemovePolicy(...)

	//e.AddRoleForUser("alice", "orderAdmin")

	//e.AddPolicy("orderAdmin", "orderList", "read")
	//e.AddPolicy("orderAdmin", "orderList", "write")

	// Check the permission.
	//b, _ := e.Enforce("bot933138", "order", "read")
	//logs.Debug(b)
	//b, _ = e.Enforce("bot933138", "order", "write")
	//logs.Debug(b)

	// Save the policy back to DB.
	//e.SavePolicy()

}

// 插入测试数据，一百万条数据
func test2(e *casbin.Enforcer) {

	// 声明业务协程处理业务
	business := func(_e *casbin.Enforcer, c chan int) {
		for {
			select {
			case i, ok := <-c:
				if !ok {
					return
				}
				// 业务
				botName := fmt.Sprintf("bot%d", i)
				logs.Debug(botName)
				_, err := _e.AddPolicy(botName, "order", "read")
				if err != nil {
					panic(err)
				}
				break
			}
		}
	}

	// 执行
	for i := 0; i < 1000000; i++ {
		c := make(chan int)
		go business(e, c)
		c <- i
		close(c)
		time.Sleep(time.Millisecond)
	}

	for {
		time.Sleep(time.Second)
	}
}

func GetEnforcer() *casbin.Enforcer {
	return e
}
