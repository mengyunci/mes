package models

import (
	"time"

	"fmt"

	"github.com/astaxie/beego/orm"
)

type Equipmentdata struct {
	CollectTime        time.Time `orm:"pk;column(CollectTime);type(datetime)"`
	DeviceID           string    `orm:"column(DeviceID);size(20)"`
	EquipmentID        string    `orm:"column(EquipmentID);size(20)"`
	Feedrate           string    `orm:"column(feedrate);size(20);null"`
	Spindlespeed       string    `orm:"column(spindlespeed);size(20);null"`
	Programnumber      string    `orm:"column(Programnumber);size(20);null"`
	Spindleload        string    `orm:"column(spindleload);size(20);null"`
	Actualfeedrate     string    `orm:"column(actualfeedrate);size(20);null"`
	Feedbeilv          string    `orm:"column(feedbeilv);size(20);null"`
	Actualspindlespeed string    `orm:"column(actualspindlespeed);size(20);null"`
	Spindlebeilv       string    `orm:"column(spindlebeilv);size(20);null"`
	Sequencenumber     string    `orm:"column(Sequencenumber);size(20);null"`
	Executingcode      string    `orm:"column(executingcode);size(20);null"`
	Xmachine           string    `orm:"column(Xmachine);size(20);null"`
	Ymachine           string    `orm:"column(Ymachine);size(20);null"`
	Zmachine           string    `orm:"column(Zmachine);size(20);null"`
	Amachine           string    `orm:"column(Amachine);size(20);null"`
	Bmachine           string    `orm:"column(Bmachine);size(20);null"`
	Xabsolute          string    `orm:"column(Xabsolute);size(20);null"`
	Yabsolute          string    `orm:"column(Yabsolute);size(20);null"`
	Zabsolute          string    `orm:"column(Zabsolute);size(20);null"`
	Aabsolute          string    `orm:"column(Aabsolute);size(20);null"`
	Babsolute          string    `orm:"column(Babsolute);size(20);null"`
	Xrelative          string    `orm:"column(Xrelative);size(20);null"`
	Yrelative          string    `orm:"column(Yrelative);size(20);null"`
	Zrelative          string    `orm:"column(Zrelative);size(20);null"`
	Arelative          string    `orm:"column(Arelative);size(20);null"`
	Brelative          string    `orm:"column(Brelative);size(20);null"`
	Emergency          string    `orm:"column(emergency);size(20);null"`
	RunModble          string    `orm:"column(runModble);size(20);null"`
	Axismove           string    `orm:"column(axismove);size(20);null"`
	Workstate          string    `orm:"column(workstate);size(20);null"`
	ALRMstate          string    `orm:"column(ALRMstate);size(20);null"`
	MaxaxisNUM         string    `orm:"column(MaxaxisNUM);size(20);null"`
	ValidaxisNUM       string    `orm:"column(validaxisNUM);size(20);null"`
}

func init() {
	orm.RegisterModel(new(Equipmentdata))
}

func GetDataEquipmentId(equipmentId string) (e []Equipmentdata, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select * from Equipmentdata where equipmentid = ? limit 1", equipmentId).QueryRows(&e)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return

}
