package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shampoo6/beemongo/models"
)

type CasbinTestController struct {
	beego.Controller
}

// @router /obj/write [get,post]
func (c *CasbinTestController) ObjWrite() {
	c.Data["json"] = models.CSuccessResponse("Obj Write Success")
	c.ServeJSON()
}

// @router /obj/write/:id string [get]
func (c *CasbinTestController) ObjWriteId() {
	s := c.Ctx.Input.Param(":id")
	c.Data["json"] = models.CSuccessResponse("Obj Write Success, id: " + s)
	c.ServeJSON()
}
