package controllers

import "hkllzh.com/easy-bill/api/models"

// BillController 账单处理 C
type BillController struct {
	EasyBillBaseController
}

// BillAddParam 账单增加的参数
type BillAddParam struct {
	Type   int     `json:"type"`
	Money  float32 `json:"money"`
	BillID int     `json:"billId"`
	UserID int     `json:"userId"`
}

// @Title 增加账单
// @Description 增加账单
// @Param body body controllers.BillAddParam true "账单信息"
// @router /add [post]
func (c *BillController) Add() {
	addParam := BillAddParam{}
	c.GetParam(&addParam)
	addParam.UserID = c.GetUserID()

	// tag := models.Tag{}
	// tag.UserID = c.GetUserID()
	// tag.Name = addParam.Name

	// orm := orm.NewOrm()
	// orm.Read(&tag, "name", "userId")

	// if 0 == tag.ID {
	// 	tag.Save()
	// 	// c.SetData(models.TrueData(tag))
	// 	c.SetData(models.NullData())
	// } else {
	// 	c.SetData(models.FalseData(models.StatusTagExist))
	// }

	c.SetData(models.TrueData(addParam))
	c.ServeJSON()
}
