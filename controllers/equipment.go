package controllers

import (
	//	"fmt"
	"mes/models"

	"github.com/astaxie/beego"
)

type EquipmentController struct {
	beego.Controller
}

func (c *EquipmentController) GetEquipmentAllState() {
	v, err := models.GetEquipmentAllState()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
