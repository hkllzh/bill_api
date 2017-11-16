package models

import (
	"github.com/astaxie/beego/orm"
)

// EasyBillBaseModel base model
type EasyBillBaseModel struct {
}

// Save 保存对象
func (m *EasyBillBaseModel) Save() {
	o := orm.NewOrm()
	o.Insert(m)
}
