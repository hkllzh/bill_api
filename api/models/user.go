package models

import "github.com/astaxie/beego/orm"

type User struct {
	ID       int    `orm:"column(id)" json:"userId"`
	Username string `orm:"size(100);column(username)" json:"username"`
	Password string `orm:"size(100);column(password)" json:"-"`
	Token    string `orm:"-" json:"token,omitempty"`
}

func (u *User) Save() {
	o := orm.NewOrm()
	o.Insert(u)
}
