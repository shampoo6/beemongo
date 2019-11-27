package filters

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/shampoo6/beemongo/casbin"
	"github.com/shampoo6/beemongo/constants"
	"github.com/shampoo6/beemongo/errors"
	"strings"
)

func mainFilter(ctx *context.Context) {
	//e := casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")
	//_, ok := ctx.Input.Session("uid").(int)
	//if !ok {
	//	ctx.Redirect(302, "/login")
	//}

	e := casbin.GetEnforcer()
	user := getUser(ctx)
	path := ctx.Request.URL.Path
	method := ctx.Request.Method
	b := e.Enforce(user, path, method)

	if !b {
		panic(errors.CError(constants.AuthError, fmt.Sprintf("%s: %s [%s]", user, path, method)))
	}

	logs.Debug("filter: %s", ctx.Request.URL.Path)
}

func getUser(ctx *context.Context) string {
	// 从不同渠道获取user
	// ctx.Input.Query("user")
	// ctx.Request.Form["user"][0]
	// ctx.Request.Header["User"][0]

	// 这里通过前端传来的user的值作为登录信息
	// 默认为未登录状态
	//user := "free"
	// 从header中获取信息
	userArr := ctx.Request.Header["User"]
	if userArr != nil && len(userArr) > 0 {
		return userArr[0]
	}
	// 从queryString中获取user
	user := ctx.Input.Query("user")
	user = strings.Trim(user, " ")
	if user != "" {
		return user
	}
	// 从form中获取user
	userArr = ctx.Request.Form["user"]
	if userArr != nil && len(userArr) > 0 {
		return userArr[0]
	}
	return "free"
}
