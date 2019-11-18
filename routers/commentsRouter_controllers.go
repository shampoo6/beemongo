package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "DeleteAll",
			Router:           `/delete`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Insert",
			Router:           `/insert`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Page",
			Router:           `/page`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shampoo6/beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Update",
			Router:           `/update`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
