package models

// User 用户实体类
type User struct {
	EasyBillBaseModel
	ID       int    `orm:"column(id)" json:"userId"`
	Username string `orm:"size(100);column(username)" json:"username"`
	Password string `orm:"size(100);column(password)" json:"-"`
	Token    string `orm:"-" json:"token,omitempty"`
}
