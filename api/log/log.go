package log

import "github.com/astaxie/beego/logs"

func init() {
	logs.SetLogger(logs.AdapterConsole)
}
