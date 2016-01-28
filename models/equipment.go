package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Equipment struct {
	BeltlineID    string    `orm:"pk;column(BeltlineID);size(20)"`
	EquipmentID   string    `orm:"column(EquipmentID);size(20)"json:"equipmentid"`
	Alarm         string    `orm:"column(alarm);size(20);null"`
	CollectTime   time.Time `orm:"column(collectTime);type(datetime);null"`
	EquipmentName string    `orm:"column(equipmentName);size(20);null"`
	EquipmentType string    `orm:"column(equipmentType);size(20);null"`
	FixLocation   string    `orm:"column(fixLocation);size(20);null"`
	Oporation     int       `orm:"column(oporation);null"`
	EquipmentData string    `orm:"column(equipmentData);size(20);null"`
}

func init() {
	orm.RegisterModel(new(Equipment))
}

func GetEquipmentId(equipmentId string) (e []Equipment, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select * from equipment where equipmentid = ?", equipmentId).QueryRows(&e)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return

}
func GetEquipmentAllState() (e []Equipment, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select * from equipment").QueryRows(&e)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return
}
