package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
)

type Module struct {
	Id          int    `orm:"column(id);auto"`
	Authorize   uint64 `orm:"column(authorize);size(1)"`
	Description string `orm:"column(description);size(255);null"`
	Name        string `orm:"column(name);size(50)"`
	Priority    int    `orm:"column(priority)"`
	ParentId    int    `orm:"column(parentId);null"`
	IconCls     string `orm:"column(iconCls);size(255);null"`
	Url         string `orm:"column(url);size(255);null"`
	Scope       string `orm:"column(scope);size(50);null"`
	PermCode    string `orm:"column(permCode);size(50);null"`
}

func (t *Module) TableName() string {
	return "module"
}

var c cache.Cache

func init() {
	c = cache.NewMemoryCache()
	orm.RegisterModel(new(Module))
}

// AddModule insert a new Module into database and returns
// last inserted Id on success.
func AddModule(m *Module) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetModuleById retrieves Module by Id. Returns error if
// Id doesn't exist
func GetModuleById(id int) (v *Module, err error) {
	o := orm.NewOrm()
	v = &Module{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllModule retrieves all Module matches certain condition. Returns empty list if
// no records exist
func GetAllModule(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Module))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Module
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateModule updates Module by Id and returns error if
// the record to be updated doesn't exist
func UpdateModuleById(m *Module) (err error) {
	o := orm.NewOrm()
	v := Module{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteModule deletes Module by Id and returns error if
// the record to be deleted doesn't exist
func DeleteModule(id int) (err error) {
	o := orm.NewOrm()
	v := Module{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Module{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetAllModuleByPriority() (ml []Module, err error) {
	// TODO need add reset cache
	if !c.IsExist("modules") {
		o := orm.NewOrm()
		num, err := o.Raw("SELECT * FROM mes.module order by priority").QueryRows(&ml)
		if err != nil {
			fmt.Println("Returned Rows Num: %s, %s", num, err)
			return nil, err
		}
		c.Put("modules", ml, 0)

	} else {
		ml = c.Get("modules").([]Module)
	}
	return

}
