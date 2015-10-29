package main

import (
	"github.com/astaxie/beego"
	_ "mes/routers"
)

func main() {
	// 运行系统
	beego.Run()
}
