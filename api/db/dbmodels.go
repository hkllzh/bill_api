package db

// User 数据库对应的用户表
type User struct {
	ID       int
	Username string `orm:"size(100)"`
	Password string `orm:"size(100)"`
}
