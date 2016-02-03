package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ShPlan struct {
	Id               int       `orm:"column(plancode);auto"json:"id"`
	Planname         string    `orm:"column(planname);size(50);null"json:"planName"`
	Starttime        time.Time `orm:"column(starttime);type(datetime);null"json:"startTime"`
	Finishtime       time.Time `orm:"column(finishtime);type(datetime);null"json:"finishTime"`
	Beltline         string    `orm:"column(beltline);size(50);null"json:"beltline`
	ProductmodelId   int       `orm:"column(productmodel_id);null"`
	Plancount        int       `orm:"column(plancount);null"json:"plancount"`
	Completecount    int       `orm:"column(completecount);null"json:"completecount"`
	Planperson       string    `orm:"column(planperson);size(50);null"json:"planperson"`
	Plandate         time.Time `orm:"column(plandate);type(datetime);null"json:"plandate"`
	Status           int       `orm:"column(status);null"json:"status"`
	Actualstarttime  time.Time `orm:"column(actualstarttime);type(datetime);null"json:"actualStartTime"`
	Actualfinishtime time.Time `orm:"column(actualfinishtime);type(datetime);null"json:"actualFinishTime"`
	Unqualify        int       `orm:"column(unqualify);null"json:"unqualify"`
}

func (t *ShPlan) TableName() string {
	return "sh_plan"
}

func init() {
	orm.RegisterModel(new(ShPlan))
}

// AddShPlan insert a new ShPlan into database and returns
// last inserted Id on success.
func AddShPlan(m *ShPlan) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetShPlanById retrieves ShPlan by Id. Returns error if
// Id doesn't exist
func GetShPlanById(id int) (v *ShPlan, err error) {
	o := orm.NewOrm()
	v = &ShPlan{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllShPlan retrieves all ShPlan matches certain condition. Returns empty list if
// no records exist
func GetAllShPlan(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ShPlan))
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

	var l []ShPlan
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

// UpdateShPlan updates ShPlan by Id and returns error if
// the record to be updated doesn't exist
func UpdateShPlanById(m *ShPlan) (err error) {
	o := orm.NewOrm()
	v := ShPlan{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteShPlan deletes ShPlan by Id and returns error if
// the record to be deleted doesn't exist
func DeleteShPlan(id int) (err error) {
	o := orm.NewOrm()
	v := ShPlan{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ShPlan{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
