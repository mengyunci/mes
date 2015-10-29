package routers

import (
	"github.com/astaxie/beego"
	"mes/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}
