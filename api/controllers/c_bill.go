package controllers

import "hkllzh.com/easy-bill/api/models"
import "github.com/astaxie/beego/logs"
import "github.com/astaxie/beego/orm"

// BillController 账单处理 C
type BillController struct {
	EasyBillBaseController
}

// BillAddParam 账单增加的参数
type BillAddParam struct {
	Type   string  `json:"type"`
	Money  float32 `json:"money"`
	TagID  []int   `json:"tagId,omitempty"`
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

	// 账单类型判断
	if addParam.Type != models.BillEarning && addParam.Type != models.BillExpense {
		c.SetData(models.FalseData(models.StatusBillSaveTypeError))
		c.ServeJSON()
		return
	}

	// golang 数组
	// slice := make([]int, 3, 5)
	// slice := make([]string, 5)
	// slice = append(slice, "1")
	// slice = append(slice, "2")
	// slice := []string{}
	// slice = append(slice, "1")
	// slice = append(slice, "x1")
	// slice = append(slice, "1d")
	// logs.Info(slice)
	// for _, str := range slice {
	// 	logs.Debug("slice->", str)
	// }

	// 可以使用的标签id，防止错误的一个校验
	canUseTagID := []int{}

	// 判断账单标签是否存在
	o := orm.NewOrm()
	for _, i := range addParam.TagID {
		t := models.Tag{}
		t.ID = i
		o.Read(&t, "id")
		if "" != t.Name {
			canUseTagID = append(canUseTagID, i)
		} else {
			logs.Error("此 tagId 不存在, tagId ->", i)
		}
	}

	// 开启数据库实务
	o.Begin()

	bill := models.Bill{}
	bill.Money = addParam.Money
	bill.UserID = c.GetUserID()
	bill.Type = addParam.Type
	_, billSaveErr := bill.Save()

	// 保存账单
	if nil != billSaveErr {
		c.SetData(models.FalseData(models.StatusBillSaveError))
		o.Rollback()
	}

	// 保存（用户 & 标签 & 账单）关系

	c.ServeJSON()
}
