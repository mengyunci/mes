package routers

import (
	"fmt"
	"mes/controllers"
	"mes/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/index", &controllers.LoginController{}, "*:Index")

	beego.Router("/equipmentMonitor/equipmentallstate", &controllers.EquipmentController{}, "*:GetEquipmentAllState")
	beego.Router("/equipmentMonitor/findByEquipmentId", &controllers.EquipmentDataController{}, "*:GetEquipmentId")

	beego.Router("/mes", &controllers.MesController{})

	beego.InsertFilter("/*", beego.BeforeRouter, func(c *context.Context) {

		_, ok := c.Input.Session("username").(string)

		if !ok && c.Input.IsAjax() {
			c.Output.JSON("relogin", true, true)

		} else if !ok && c.Request.RequestURI != "/login" {
			c.Redirect(302, "/login")
		} else {
			modules, err := models.GetAllModuleByPriority()
			if err != nil {
				fmt.Println(err)
			}
			c.Input.SetData("modules", modules)
			c.Input.SetData("url", c.Request.RequestURI)
			c.Input.SetData("username", c.Input.Session("username"))
		}
	})
}
