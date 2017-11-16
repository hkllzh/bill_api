package controllers

import (
	"hkllzh.com/easy-bill/api/models"
)

// TagController 标签处理 C
type TagController struct {
	EasyBillBaseController
}

// TagAddParam 标签增加的参数
type TagAddParam struct {
	Name string `json:"name"`
}

// @Title 增加标签
// @Description 增加标签
// @Param body body models.User true "用户信息"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /add [post]
func (c *TagController) Add() {
	addParam := TagAddParam{}
	c.GetParam(&addParam)
	//	c.Data["json"] = models.FalseData(models.StatusLoginFailed)
	c.SetData(models.TrueData(addParam))
	c.ServeJSON()
}
