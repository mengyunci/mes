package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplNames = "login.html"
}

func (this *MainController) LoginGet() {
	this.TplNames = "login.tpl"
}
