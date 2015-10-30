package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


func (this *MainController) Prepare() {
	this.TplNames = "login.html"
	username := this.GetSession("username")
	fmt.Println(username)
	if username != nil  {
		this.TplNames = "index.html"
	} else {
		this.Redirect("/index" , 302)
	}
	this.StopRun()
}


