package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {

	dns := beego.AppConfig.String("dns")
	if err := orm.RegisterDataBase("default", "mysql", dns); err != nil {
		panic(err)
	}

}
