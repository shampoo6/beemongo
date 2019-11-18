package conf

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/shampoo6/beemongo/errors"
	"github.com/shampoo6/beemongo/models"
	"reflect"
)

func overrideRecoverFunc() {
	recoverFunc := beego.BConfig.RecoverFunc
	newFunc := func(c *context.Context) {
		defer recoverFunc(c)
		if err := recover(); err != nil {
			// 捕获到异常后自己先处理，再抛出去给 beego 处理
			// 设置响应码为200
			c.Output.Status = 200
			// 判断是业务异常还是系统异常
			var responseStr string
			if reflect.TypeOf(err).Name() == errors.BusinessErrorName {
				_ = c.Output.JSON(models.CBusinessResponse(err.(errors.BusinessError)), beego.BConfig.RunMode != beego.PROD, false)
			} else {
				_ = c.Output.JSON(models.CExceptionResponse(fmt.Sprintf("Handler crashed with error: %s", err)), beego.BConfig.RunMode != beego.PROD, false)
			}
			c.ResponseWriter.Header().Set("Content-Type", "application/json;charset=utf-8")
			_, _ = c.ResponseWriter.Write([]byte(responseStr))
			panic(err)
		}
	}
	beego.BConfig.RecoverFunc = newFunc
}
