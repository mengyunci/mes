package routers

import (
	"fmt"
	"mes/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/index", &controllers.LoginController{}, "*:Index")

	beego.Router("/equipmentMonitor/equipmentallstate", &controllers.EquipmentController{}, "*:GetEquipmentAllState")
	beego.Router("/equipmentMonitor/findByEquipmentId", &controllers.EquipmentDataController{}, "*:GetEquipmentId")

	beego.InsertFilter("/*", beego.BeforeRouter, func(c *context.Context) {

		_, ok := c.Input.Session("username").(string)

		if !ok && c.Input.IsAjax() {
			fmt.Println("ajax")
			c.Output.JSON("relogin", true, true)

		}

		if !ok && c.Request.RequestURI != "/login" {
			fmt.Println("normal")
			c.Redirect(302, "/login")
		}
	})

	beego.InsertFilter("/*", beego.BeforeExec, func(c *context.Context) {

		c.Input.SetData("username", c.Input.Session("username"))
	})
}
