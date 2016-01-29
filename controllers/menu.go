package controllers

import (
	"encoding/json"
	"errors"
	"mes/models"
	"strconv"
	"strings"

	"fmt"
	"github.com/astaxie/beego"
)

// oprations for Menu
type MenuController struct {
	beego.Controller
}

func (c *MenuController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Menu
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 201 {int} models.Menu
// @Failure 403 body is empty
// @router / [post]
func (c *MenuController) Post() {
	var v models.Menu
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddMenu(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get
// @Description get Menu by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Menu
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MenuController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetMenuById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Menu
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Menu
// @Failure 403
// @router / [get]
func (c *MenuController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllMenu(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Menu
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 200 {object} models.Menu
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MenuController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Menu{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateMenuById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Menu
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MenuController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteMenu(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *MenuController) LoadByModuleId() {
	mid := c.Ctx.Input.Param(":moduleid")

	query := make(map[string]string)
	query["ModuId"] = mid
	q := make([]string, 0)
	menus, err := models.GetAllMenu(query, q, q, q, 0, -1)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		rMenu := make([]models.Menu, 0)
		pMenu := make(map[int]models.Menu)
		for _, v := range menus {
			vt, _ := v.(models.Menu)
			if vt.ParentId == 0 {
				pMenu[vt.Id] = vt
			}
		}
		for _, v := range menus {
			vt, _ := v.(models.Menu)
			if vt.ParentId != 0 {
				t := pMenu[vt.ParentId]
				t.Children = append(t.Children, vt)
				pMenu[vt.ParentId] = t
			}
		}

		for _, v := range pMenu {
			rMenu = append(rMenu, v)
		}

		fmt.Println()
		c.Data["json"] = rMenu
	}
	c.ServeJSON()
}
