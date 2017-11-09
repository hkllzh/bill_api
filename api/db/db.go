package db

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	fmt.Println("package db func init() {")
	fmt.Println("连接数据库")

	// set default database
	orm.RegisterDataBase("default", "mysql", "easy_bill:pwd_easy_bill@tcp(121.196.217.69:3306)/easy_bill?charset=utf8", 30)

	// register model
	orm.RegisterModelWithPrefix("t_", new(User))

	orm.Debug = true

	// create table
	// orm.RunSyncdb("default", false, true)
}
