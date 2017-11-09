package db

import "github.com/astaxie/beego/orm"

// User 数据库对应的用户表
type User struct {
	ID       int    `orm:"column(id)"`
	Username string `orm:"size(100)"`
	Password string `orm:"size(100)"`
	Age      int
}

func (u *User) Save() {
	o := orm.NewOrm()
	o.Insert(u)
}
