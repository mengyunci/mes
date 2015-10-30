package routers

import (
	"github.com/astaxie/beego"
	"mes/controllers"
	"github.com/astaxie/beego/context"
	"fmt"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/index", &controllers.LoginController{}, "*:Index")

	beego.InsertFilter("/*",beego.BeforeRouter, func(c *context.Context) {
		fmt.Println(c.Request.RequestURI)
		_,ok := c.Input.Session("username").(string)
		if !ok && c.Request.RequestURI != "/login"  {
			c.Redirect(302, "/login")
		}
	})

	beego.InsertFilter("/*",beego.BeforeExec, func(c *context.Context) {
		c.Input.Data["username"] = c.Input.Session("username")
	})
}
