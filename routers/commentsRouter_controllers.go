package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Insert",
			Router:           `/insert`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Page",
			Router:           `/page`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beemongo/controllers:UserController"] = append(beego.GlobalControllerRouter["beemongo/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Update",
			Router:           `/update`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
