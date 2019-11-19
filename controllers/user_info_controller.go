package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shampoo6/beemongo/errors"
	"github.com/shampoo6/beemongo/models"
	"github.com/shampoo6/beemongo/models/dto"
	"github.com/shampoo6/beemongo/service/user_info"
)

type UserInfoController struct {
	beego.Controller
}

// @router /insert [post,get]
func (c *UserInfoController) Insert() {
	dto := dto.UserInfoDto{}
	if err := c.ParseForm(&dto); err != nil {
		panic(err)
	}
	valid := dto.Validation()
	b, err := valid.Valid(&dto)
	if err != nil {
		panic(err)
	}
	if !b {
		panic(errors.CParamError(valid.Errors))
	}
	user_info.Insert(&dto)
	c.Data["json"] = models.CSuccessResponse("Ok")
	c.ServeJSON()
}
