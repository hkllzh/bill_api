package controllers

import (
	"encoding/json"

	"hkllzh.com/easy-bill/api/models"

	"fmt"
	"github.com/astaxie/beego/orm"
	"hkllzh.com/easy-bill/api/cache"
)

// Operations about Users
type UserController struct {
	EasyBillBaseController
}

// @Title 注册用户
// @Description 注册用户
// @Param body body models.User true "用户信息"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /register [post]
func (u *UserController) Register() {

	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	o := orm.NewOrm()
	o.Read(&user, "username")

	if 0 == user.ID {

		user.Save()

		user.Token = cache.GetToken()
		cache.PutUserToken(user)

		u.Data["json"] = models.TrueData(user)
	} else {
		u.Data["json"] = models.FalseData(1000, "账号已经已经注册")
	}

	fmt.Println(user)

	u.ServeJSON()

}

// @Title 用户登录
// @Description 用户登录
// @Param body body models.User true "用户信息"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /login [post]
func (u *UserController) Login() {

	fmt.Println("UserController -> Login")
	fmt.Println(u.Ctx.Request.Header)
}

//// @Title CreateUser
//// @Description create users
//// @Param	body		body 	models.User	true		"body for user content"
//// @Success 200 {int} models.User.Id
//// @Failure 403 body is empty
//// @router / [post]
//func (u *UserController) Post() {
//	var user models.User
//	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//	uid := models.AddUser(user)
//	u.Data["json"] = map[string]string{"uid": uid}
//	u.ServeJSON()
//}
//
//// @Title GetAll
//// @Description get all Users
//// @Success 200 {object} models.User
//// @router / [get]
//func (u *UserController) GetAll() {
//	users := models.GetAllUsers()
//	u.Data["json"] = users
//	u.ServeJSON()
//}
//
//// @Title Get
//// @Description get user by uid
//// @Param	uid		path 	string	true		"The key for staticblock"
//// @Success 200 {object} models.User
//// @Failure 403 :uid is empty
//// @router /:uid [get]
//func (u *UserController) Get() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		user, err := models.GetUser(uid)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = user
//		}
//	}
//	u.ServeJSON()
//}
//
//// @Title Update
//// @Description update the user
//// @Param	uid		path 	string	true		"The uid you want to update"
//// @Param	body		body 	models.User	true		"body for user content"
//// @Success 200 {object} models.User
//// @Failure 403 :uid is not int
//// @router /:uid [put]
//func (u *UserController) Put() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		var user models.User
//		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//		uu, err := models.UpdateUser(uid, &user)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = uu
//		}
//	}
//	u.ServeJSON()
//}
//
//// @Title Delete
//// @Description delete the user
//// @Param	uid		path 	string	true		"The uid you want to delete"
//// @Success 200 {string} delete success!
//// @Failure 403 uid is empty
//// @router /:uid [delete]
//func (u *UserController) Delete() {
//	uid := u.GetString(":uid")
//	models.DeleteUser(uid)
//	u.Data["json"] = "delete success!"
//	u.ServeJSON()
//}
//
//// @Title logout
//// @Description Logs out current logged in user session
//// @Success 200 {string} logout success
//// @router /logout [get]
//func (u *UserController) Logout() {
//	u.Data["json"] = "logout success"
//	u.ServeJSON()
//}
