package models

//type User struct {
//	ID       string
//	Username string
//	Password string
//}

const (
	StatusSuccess      = 0
	StatusUnknownError = 10000
	StatusAuthFailed   = 10001
	StatusAccountExist = 10002
	StatusLoginFailed  = 10003
)

var errMessageCH = map[int]string{
	StatusSuccess:      "",
	StatusUnknownError: "未知错误",
	StatusAuthFailed:   "请求验证失败",
	StatusAccountExist: "账号已经存在",
	StatusLoginFailed:  "账号或密码错误",
}

type ReturnData struct {
	Code   int         `json:"code"`           // 请求代码
	ErrMsg string      `json:"errMsg"`         // 错误原因
	Data   interface{} `json:"data,omitempty"` // 返回数据
}

func TrueData(data interface{}) ReturnData {
	var rData ReturnData
	rData.Code = StatusSuccess
	rData.Data = data
	rData.ErrMsg = errMessageCH[StatusSuccess]
	return rData
}

func FalseData(code int) ReturnData {
	var rData ReturnData
	rData.Code = code
	rData.Data = nil
	rData.ErrMsg = errMessageCH[code]
	return rData
}
