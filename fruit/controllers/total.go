package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type TotalController struct {
	beego.Controller
}

// @router /total/:name [get]
func (c *TotalController) Get() {

	name := c.Ctx.Input.Param(":name")

	fmt.Println("name: " + name)

	return
}
