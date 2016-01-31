package controllers

import (
	//	"fmt"

	"fmt"
	"github.com/astaxie/beego"
	"mes/models"
	"strconv"
)

type MesController struct {
	beego.Controller
}

func (this *MesController) Get() {

	ms, err := models.GetAllModuleByPriority()
	if err != nil {
		fmt.Println(err)
	}
	var moudleID int
	for _, v := range ms {
		if v.Url == "/mes" {
			moudleID = v.Id
			break
		}
	}

	query := make(map[string]string)
	query["ModuId"] = strconv.Itoa(moudleID)
	q := make([]string, 0)
	menus, err := models.GetAllMenu(query, q, []string{"desc", "desc"}, []string{"ParentId", "ModuId"}, 0, -1)
	if err != nil {
		fmt.Println(err)
	}
	this.Data["menus"] = menus

	this.Data["title"] = "mes"
	this.TplName = "mes.html"
}
