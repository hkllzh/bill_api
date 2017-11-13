//
// 用户相关的缓存模块代码
//
package cache

import (
	"github.com/astaxie/beego/logs"
	"strconv"
	"hkllzh.com/easy-bill/api/models"
	"time"
	"fmt"
)

const UserToken = "user_token_"

func PutUserToken(user models.User) {
	logs.Debug("PutUserToken->", user)
	logs.Debug("PutUserToken->", UserToken+strconv.Itoa(user.ID))
	logs.Debug("PutUserToken->", user.Token)
	// 一天的缓存时间
	err := fileCache.Put(UserToken+strconv.Itoa(user.ID), user.Token, 24*60*60*time.Second)
	if nil != err {
		logs.Error("缓存出错 user->", user)
	}
}

func GetUserToken(user models.User) string {
	if fileCache.IsExist(UserToken + strconv.Itoa(user.ID)) {
		return fmt.Sprint(fileCache.Get(UserToken + strconv.Itoa(user.ID)))
	} else {
		return ""
	}
}
