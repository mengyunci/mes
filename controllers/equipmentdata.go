package controllers

import (
	//	"fmt"
	"fmt"
	"mes/models"

	"github.com/astaxie/beego"
)

type EquipmentDataController struct {
	beego.Controller
}

func (c *EquipmentDataController) GetEquipmentId() {
	idStr := c.GetString("equipmentId")
	//	idStr := c.Ctx.Input.Param(":equipmentId")
	fmt.Println(idStr + "----------")
	v, err := models.GetDataEquipmentId(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()

}
