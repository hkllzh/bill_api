package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"hkllzh.com/easy-bill/api/cache"
	"hkllzh.com/easy-bill/api/models"
)

type EasyBillBaseController struct {
	beego.Controller
}

func (c *EasyBillBaseController) Prepare() {

	// 需要鉴权 && 不通过
	if needAuthUrlFilter(c) && !tokenAuth(c) {
		c.Data["json"] = models.FalseData(models.StatusAuthFailed)
		c.ServeJSON()
		c.StopRun()
	}
}

// SetData 设置返回值
func (c *EasyBillBaseController) SetData(data models.ReturnData) {
	c.Data["json"] = data
}

func (c *EasyBillBaseController) GetParam(param interface{}) {
	json.Unmarshal(c.Ctx.Input.RequestBody, param)
}

// 判断是否需要鉴权
func needAuthUrlFilter(c *EasyBillBaseController) bool {
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

	logs.Debug(header)

	token := header.Get("Token")

	if "" == token {
		return false
	}

	tokens := strings.Split(token, ",")
	id, err := strconv.Atoi(tokens[0])
	if nil != err {
		logs.Error("tokenAuth", err)
		return false
	}
	u := models.User{ID: id}
	cacheToken := cache.GetUserToken(u)

	if cacheToken != tokens[1] {
		return false
	}

	return true
}
