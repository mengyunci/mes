package controllers

import (
	//	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	this.SetSession("username", "admin")
	this.Redirect("/index", 302)
	//	username := this.GetString("username")
	//	password := this.GetString("password")
	//	if username == "admin" && password == "admin" {
	//		this.SetSession("username", username)
	//		this.Redirect("/index", 302)
	//	} else {
	//		this.Redirect("/login", 302)
	//	}
}

func (this *LoginController) Index() {
	this.Data["title"] = "欢迎"
	this.TplName = "index.html"
	//	username := this.GetSession("username")
	//	//	fmt.Println(username)

	//	if username != nil {
	//		this.Data["title"] = "欢迎"
	//		this.TplName = "index.html"
	//	} else {
	//		this.Redirect("/login", 302)
	//	}
}
