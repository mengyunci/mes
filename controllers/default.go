package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {
	this.TplName = "login.html"
	username := this.GetSession("username")
	fmt.Println(username)
	if username != nil {
		this.TplName = "index.html"
	} else {
		this.Redirect("/index", 302)
	}
	this.StopRun()
}
