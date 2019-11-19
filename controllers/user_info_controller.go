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
	_dto := getParam(c)
	c.Data["json"] = models.CSuccessResponse(user_info.Insert(&_dto))
	c.ServeJSON()
}

func getParam(c *UserInfoController) dto.UserInfoDto {
	_dto := dto.UserInfoDto{}
	if err := c.ParseForm(&_dto); err != nil {
		panic(err)
	}
	valid := _dto.Validation()
	b, err := valid.Valid(&_dto)
	if err != nil {
		panic(err)
	}
	if !b {
		panic(errors.CParamError(valid.Errors))
	}
	return _dto
}
