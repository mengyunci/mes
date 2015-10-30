package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	if username == "admin" && password == "admin" {
		this.SetSession("username",username)
		this.Redirect("/index",302)
	} else {
		this.Redirect("/login",302)
	}
}

func (this *LoginController) Index() {
	username := this.GetSession("username")
	fmt.Println(username)

	if  username != nil {
		this.Data["title"] = "欢迎"
		this.TplNames = "index.html"
	} else {
		this.Redirect("/login",302)
	}
}
