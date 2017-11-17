package models

import "github.com/astaxie/beego/orm"

// Bill 账单实体类
type Bill struct {
	EasyBillBaseModel
	ID     int     `orm:"column(id)" json:"billId"`
	UserID int     `orm:"column(user_id)" json:"-"`     // 用户id
	TagID  int     `orm:"column(tag_id)" json:"-"`      // 标签id
	Money  float32 `orm:"column(money)" json:"money"`   // 账单金额
	Type   int     `orm:"column(type)" json:"type"`     // 类型 1 收入、2 支出
	Remark string  `orm:"column(remark)" json:"remark"` // 备注
}

// Save 保存对象
func (m *Bill) Save() {
	o := orm.NewOrm()
	o.Insert(m)
}
