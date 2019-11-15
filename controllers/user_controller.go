package controllers

import (
	"beemongo/models"
	"beemongo/service/user"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

// @router /insert [get,post]
func (c *UserController) Insert() {
	result := map[string]interface{}{"result": "ok"}
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
		result["result"] = "not ok"
		for _, err := range valid.Errors {
			logs.Error(err.Key, err.Message)
			result[err.Key] = err.Message
		}
	} else {
		// insert 数据进数据库
		user := user_service.Insert(&u)
		result["data"] = user
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /update [get,post]
func (c *UserController) Update() {
	result := map[string]interface{}{"result": "ok"}
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
		result["result"] = "not ok"
		for _, err := range valid.Errors {
			logs.Error(err.Key, err.Message)
			result[err.Key] = err.Message
		}
	} else {
		// insert 数据进数据库
		user := user_service.Update(&u)
		result["data"] = user
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *UserController) page() {
	page := models.Page{}
	c.ParseForm(&page)
}
