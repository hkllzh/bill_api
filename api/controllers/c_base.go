package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"hkllzh.com/easy-bill/api/cache"
	"hkllzh.com/easy-bill/api/models"
	"github.com/astaxie/beego/orm"
)

type EasyBillBaseController struct {
	beego.Controller
	ebOrm orm.Ormer
}

func (c *EasyBillBaseController) Prepare() {

	// 需要鉴权 && 不通过
	if needAuthURLFilter(c) && !tokenAuth(c) {
		c.Data["json"] = models.FalseData(models.StatusAuthFailed)
		c.ServeJSON()
		c.StopRun()
	}

	// create new orm
	c.ebOrm = orm.NewOrm()
}

// GetUserID 获取用户id
func (c *EasyBillBaseController) GetUserID() int {
	header := c.Ctx.Request.Header
	token := header.Get("Token")
	tokens := strings.Split(token, ",")
	id, err := strconv.Atoi(tokens[0])
	if nil != err {
		logs.Error("tokenAuth", err)
		return 0
	}
	return id
}

// SetData 设置返回值
func (c *EasyBillBaseController) SetData(data models.ReturnData) {
	c.Data["json"] = data
}

// GetParam 获取参数
func (c *EasyBillBaseController) GetParam(param interface{}) {
	json.Unmarshal(c.Ctx.Input.RequestBody, param)
}

// needAuthURLFilter 判断是否需要鉴权
func needAuthURLFilter(c *EasyBillBaseController) bool {
	url := c.Ctx.Request.URL.Path
	notFilter := []string{
		"/v1/user/register",
		"/v1/user/login",
	}
	need := true
	for i := 0; i < len(notFilter); i++ {
		if strings.HasPrefix(url, notFilter[i]) {
			need = false
			break
		}
	}
	return need
}

// tokenAuth 验证token
func tokenAuth(c *EasyBillBaseController) bool {
	// token 验证

	header := c.Ctx.Request.Header
	token := header.Get("Token")

	logs.Debug("token -> ", token)

	if "" == token {
		return false
	}

	tokens := strings.Split(token, ",")
	logs.Debug("tokens -> ", tokens)
	id, err := strconv.Atoi(tokens[0])
	logs.Debug("tokens ,id -> ", id)
	if nil != err {
		logs.Error("tokenAuth", err)
		return false
	}
	u := models.User{ID: id}
	cacheToken := cache.GetUserToken(u)
	logs.Debug("tokens ,cacheToken -> ", cacheToken)

	if cacheToken != tokens[1] {
		return false
	}

	return true
}
