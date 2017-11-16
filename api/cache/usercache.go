//
// 用户相关的缓存模块代码
//
package cache

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"hkllzh.com/easy-bill/api/models"
)

// UserTokenPrefix 用户token前缀
const UserTokenPrefix = "user_token_"

// UserTokenTimeout 用户token超时时间
const UserTokenTimeout = 24 * 60 * 60 * time.Second //time.Duration

// SetUserToken 设置用户token
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

// GetUserToken 获取用户token
func GetUserToken(user models.User) string {
	key := UserTokenPrefix + strconv.Itoa(user.ID)
	return getValue(key)
}
