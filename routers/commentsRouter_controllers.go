package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:CasbinTestController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:CasbinTestController"],
		beego.ControllerComments{
			Method:           "ObjWrite",
			Router:           `/obj/write`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:CasbinTestController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:CasbinTestController"],
		beego.ControllerComments{
			Method:           "ObjWriteId",
			Router:           `/obj/write/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "DeleteAll",
			Router:           `/delete`,
			AllowHTTPMethods: []string{"get", "post", "options"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Insert",
			Router:           `/insert`,
			AllowHTTPMethods: []string{"get", "post", "options"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Page",
			Router:           `/page`,
			AllowHTTPMethods: []string{"get", "post", "options"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Update",
			Router:           `/update`,
			AllowHTTPMethods: []string{"get", "post", "options"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserInfoController"],
		beego.ControllerComments{
			Method:           "Insert",
			Router:           `/insert`,
			AllowHTTPMethods: []string{"post", "get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
