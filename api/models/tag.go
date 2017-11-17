package models

import "github.com/astaxie/beego/orm"

// Tag 标签实体类
type Tag struct {
	EasyBillBaseModel
	ID        int    `orm:"column(id)" json:"tagId"`
	UserID    int    `orm:"column(userId)" json:"-"`             // 用户id
	Name      string `orm:"size(100);column(name)" json:"name"`  // 标签名字
	BillCount int    `orm:"column(bill_count)" json:"billCount"` // 对应的账单数量
}

// Save 保存对象
func (m *Tag) Save() {
	o := orm.NewOrm()
	o.Insert(m)
}
