package controllers

import (
	"github.com/astaxie/beego/orm"
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
// @Param body body controllers.TagAddParam true "用户信息"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /add [post]
func (c *TagController) Add() {
	addParam := TagAddParam{}
	c.GetParam(&addParam)

	tag := models.Tag{}
	tag.UserID = c.GetUserID()
	tag.Name = addParam.Name

	orm := orm.NewOrm()
	orm.Read(&tag, "name", "userId")

	if 0 == tag.ID {
		tag.Save()
		// c.SetData(models.TrueData(tag))
		c.SetData(models.NullData())
	} else {
		c.SetData(models.FalseData(models.StatusTagExist))
	}

	c.ServeJSON()
}
