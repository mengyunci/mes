package controllers

import (
	//	"fmt"

	"github.com/astaxie/beego"
)

type MesController struct {
	beego.Controller
}

func (this *MesController) Get() {
	this.Data["title"] = "mes"
	this.TplName = "mes.html"
}
