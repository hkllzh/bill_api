//
// 用户相关的缓存模块代码
//
package cache

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"hkllzh.com/easy-bill/api/models"
)

const UserTokenPrefix = "user_token_"
const UserTokenTimeout = 24 * 60 * 60 * time.Second //time.Duration

func SetUserToken(user models.User) {

	key := UserTokenPrefix + strconv.Itoa(user.ID)

	if fileCache.IsExist(key) {
		fileCache.Delete(key)
	}

	// 一天的缓存时间
	err := fileCache.Put(key, user.Token, UserTokenTimeout)
	if nil != err {
		logs.Error("缓存出错 user->", user)
	}
}

func GetUserToken(user models.User) string {
	if fileCache.IsExist(UserTokenPrefix + strconv.Itoa(user.ID)) {
		return fmt.Sprint(fileCache.Get(UserTokenPrefix + strconv.Itoa(user.ID)))
	} else {
		return ""
	}
}
