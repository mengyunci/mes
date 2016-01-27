package main

import (
	"mes/models"
	_ "mes/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 运行系统
	models.Init()
	beego.Run()
}
