package routers

import (
	"github.com/astaxie/beego"
	"github.com/shampoo6/beemongo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/user",
		beego.NSInclude(&controllers.UserController{}),
	)
	ns2 := beego.NewNamespace("/userInfo",
		beego.NSInclude(&controllers.UserInfoController{}),
	)
	ns3 := beego.NewNamespace("/casbin",
		beego.NSInclude(&controllers.CasbinTestController{}),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(ns2)
	beego.AddNamespace(ns3)
}
