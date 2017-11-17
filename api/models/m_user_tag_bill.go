package models

import "github.com/astaxie/beego/orm"

// UserTagBill 用户实体类
type UserTagBill struct {
	ID     int `orm:"column(id)"`
	UserID int `orm:"column(user_id)"`
	TagID  int `orm:"column(tag_id)"`
	BillID int `orm:"column(bill_id)"`
}

// Save 保存对象
func (m *UserTagBill) Save() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}
