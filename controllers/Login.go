package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
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
	fmt.Println(username)
	fmt.Println(password)
	if username == "admin" && password == "admin" {
		this.TplNames = "index.html"
	}
	this.TplNames = "login.html"
}
