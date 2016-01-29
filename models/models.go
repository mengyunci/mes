package models

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {

	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	dsn := iniconf.String("dns")
	orm.RegisterDataBase("default", "mysql", dsn)

}
