package controllers

import (
	"beemongo/errors"
	"beemongo/models"
	"beemongo/service/user"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @router /insert [get,post]
func (c *UserController) Insert() {
	u := models.UserDto{}
	if err := c.ParseForm(&u); err != nil {
		panic(err)
	}
	valid := u.Validation()
	b, err := valid.Valid(&u)
	if err != nil {
		panic(err)
	}
	if !b {
		panic(errors.CParamError(valid.Errors))
	} else {
		user := user_service.Insert(&u)
		c.Data["json"] = models.CSuccessResponse(*user)
	}
	c.ServeJSON()
}

// @router /update [get,post]
func (c *UserController) Update() {
	u := models.UserDto{}
	if err := c.ParseForm(&u); err != nil {
		panic(err)
	}
	valid := u.Validation()
	valid.Required(u.Id, "Id").Message("id不能为空")
	b, err := valid.Valid(&u)
	if err != nil {
		panic(err)
	}
	if !b {
		panic(errors.CParamError(valid.Errors))
	} else {
		user := user_service.Update(&u)
		c.Data["json"] = models.CSuccessResponse(*user)
	}
	c.ServeJSON()
}

func (c *UserController) page() {
	page := models.Page{}
	c.ParseForm(&page)
}
