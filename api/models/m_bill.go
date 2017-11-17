package models

import "github.com/astaxie/beego/orm"

// BillEarning 收入
const BillEarning = "earning"

// BillExpense 支出
const BillExpense = "expense"

// Bill 账单实体类
type Bill struct {
	EasyBillBaseModel
	ID     int     `orm:"column(id)" json:"billId"`
	UserID int     `orm:"column(user_id)" json:"-"`     // 用户id
	Money  float32 `orm:"column(money)" json:"money"`   // 账单金额
	Type   string  `orm:"column(type)" json:"type"`     // 类型 earning 收入、expense 支出
	Remark string  `orm:"column(remark)" json:"remark"` // 备注
}

// Save 保存对象
func (m *Bill) Save() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}
