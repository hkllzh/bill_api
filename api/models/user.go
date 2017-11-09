package models

import "github.com/astaxie/beego/orm"

type User struct {
	ID       int    `orm:"column(id)" json:"-"`
	Username string `orm:"size(100)" json:"username"`
	Password string `orm:"size(100)" json:"-"`
	Token    string `orm:"-" json:"token,omitempty"`
}

func (u *User) Save() {
	o := orm.NewOrm()
	o.Insert(u)
}
