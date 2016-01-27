package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {

	dsn := "root:mysql@tcp(192.168.0.123:3306)/mes?charset=utf8"

	orm.RegisterDataBase("default", "mysql", dsn)

}
