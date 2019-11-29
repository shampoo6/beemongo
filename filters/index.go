package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

func InitFilters() {
	beego.InsertFilter("*", beego.BeforeRouter, beforeRouterFilter)
}

func beforeRouterFilter(ctx *context.Context) {
	allow := cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods: []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: []string{"*"},
		//公开的HTTP标头列表
		ExposeHeaders: []string{"Access-Control-Allow-Origin", "Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	})
	allow(ctx)
	mainFilter(ctx)
}
