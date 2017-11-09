package models

//type User struct {
//	ID       string
//	Username string
//	Password string
//}

type ReturnData struct {
	Code   int         `json:"code"`           // 请求代码
	ErrMsg string      `json:"errMsg"`         // 错误原因
	Data   interface{} `json:"data,omitempty"` // 返回数据
}

func TrueData(data interface{}) ReturnData {
	var rData ReturnData
	rData.Code = 0
	rData.Data = data
	rData.ErrMsg = ""
	return rData
}

func FalseData(code int, errMsg string) ReturnData {
	var rData ReturnData
	rData.Code = code
	rData.Data = nil
	rData.ErrMsg = errMsg
	return rData
}
