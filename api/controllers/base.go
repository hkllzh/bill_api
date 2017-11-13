package controllers

import (
	"github.com/astaxie/beego"
	"hkllzh.com/easy-bill/api/models"
	"github.com/astaxie/beego/logs"
	"strings"
	"hkllzh.com/easy-bill/api/cache"
	"strconv"
)

type EasyBillBaseController struct {
	beego.Controller
}

func (c *EasyBillBaseController) Prepare() {

	// 需要鉴权 && 不通过
	if needAuthUrlFilter(c) && !tokenAuth(c) {
		c.Data["json"] = models.FalseData(10001, "验证不通过")
		c.ServeJSON()
		c.StopRun()
	}
}

// 判断是否需要鉴权
func needAuthUrlFilter(c *EasyBillBaseController) bool {
	url := c.Ctx.Request.URL.Path
	notFilter := []string{
		"/v1/user/register",
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

	logs.Debug("token->", "-", token, "-")

	tokens := strings.Split(token, ",")

	logs.Debug("token->", "-", tokens[0], "-")
	logs.Debug("token->", "-", tokens[1], "-")

	id, err := strconv.Atoi(tokens[0])
	if nil != err {
		logs.Error("tokenAuth", err)
		return false
	}
	u := models.User{ID: id}
	cacheToken := cache.GetUserToken(u)
	logs.Debug("cacheToken->", "-", cacheToken, "-")

	if cacheToken != tokens[1] {
		return false
	}

	return true
}
