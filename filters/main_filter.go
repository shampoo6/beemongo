package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func mainFilter(ctx *context.Context) {
	//e := casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")
	//_, ok := ctx.Input.Session("uid").(int)
	//if !ok {
	//	ctx.Redirect(302, "/login")
	//}
	logs.Debug("filter: %s", ctx.Request.RequestURI)
}
