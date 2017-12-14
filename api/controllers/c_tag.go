package controllers

import (
	"github.com/astaxie/beego/orm"
	"hkllzh.com/easy-bill/api/models"
	"github.com/astaxie/beego/logs"
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
// @Param body body controllers.TagAddParam true "用户信息"
// @router /add [post]
func (c *TagController) Add() {
	addParam := TagAddParam{}
	c.GetParam(&addParam)

	tag := models.Tag{}
	tag.UserID = c.GetUserID()
	tag.Name = addParam.Name

	o := orm.NewOrm()
	o.Read(&tag, "name", "userId")

	if 0 == tag.ID {
		tag.Save()
		// c.SetData(models.TrueData(tag))
		c.SetData(models.NullData())
	} else {
		c.SetData(models.FalseData(models.StatusTagExist))
	}

	c.ServeJSON()
}

// @Title 标签列表
// @Description 标签列表
// @router /list [post]
func (c *TagController) List() {

	var ts []*models.Tag

	c.ebOrm.QueryTable("t_tag").Filter("userId", c.GetUserID()).All(&ts)
	logs.Debug("ts -> ", ts)

	c.SetData(models.TrueData(ts))
	c.ServeJSON()
}
