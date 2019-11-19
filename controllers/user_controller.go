package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/shampoo6/beemongo/constants"
	"github.com/shampoo6/beemongo/errors"
	"github.com/shampoo6/beemongo/models"
	"github.com/shampoo6/beemongo/models/dto"
	"github.com/shampoo6/beemongo/service/user"
	"regexp"
)

type UserController struct {
	beego.Controller
}

// @router /insert [get,post]
func (c *UserController) Insert() {
	u := dto.UserDto{}
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
	u := dto.UserDto{}
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
	}
	user := user_service.Update(&u)
	c.Data["json"] = models.CSuccessResponse(*user)
	c.ServeJSON()
}

// @router /delete [get,post]
func (c *UserController) DeleteAll() {
	ids := c.GetStrings("ids")
	if len(ids) == 0 {
		msg := map[string]string{"ids": "id列表不能为空"}
		panic(errors.BusinessError{Msg: constants.ParamError.Remark, Status: constants.ParamError, Content: msg})
	}
	c.Data["json"] = models.CSuccessResponse(user_service.DeleteAll(ids))
	c.ServeJSON()
}

// @router /page [get,post]
func (c *UserController) Page() {
	page := models.Page{}
	if err := c.ParseForm(&page); err != nil {
		panic(err)
	}
	dto := dto.UserDto{}
	ptr := &dto
	if err := c.ParseForm(ptr); err != nil {
		panic(err)
	}
	reg, _ := regexp.Compile("^(Male|Female)?$")
	valid := validation.Validation{}
	valid.Match(ptr.Sex, reg, "Sex").Message("性别必须为 Male 或 Female")
	b, err := valid.Valid(ptr)
	if err != nil {
		panic(err)
	}
	if !b {
		panic(errors.CParamError(valid.Errors))
	} else {
		c.Data["json"] = models.CSuccessResponse(user_service.Page(&page, &dto))
	}
	c.ServeJSON()
}
