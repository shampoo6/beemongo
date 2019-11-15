package routers

import (
	"beemongo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/user",
		beego.NSInclude(&controllers.UserController{}),
	)
	beego.AddNamespace(ns)
}
