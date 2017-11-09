package cache

import (
	"github.com/astaxie/beego/cache"
	"hkllzh.com/easy-bill/api/models"
	"strconv"
	"fmt"
)

var fileCache cache.Cache

const USER_TOKEN = "user_token_"

func init() {
	fmt.Println("缓存初始化")
	bm, _ := cache.NewCache("file", `{"CachePath":"./file_cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	fileCache = bm
}

func PutUserToken(user models.User) {
	fileCache.Put(USER_TOKEN+strconv.Itoa(user.ID), user.Token, 10000)
}

func GetUserToken(user models.User) string {
	if fileCache.IsExist(USER_TOKEN + strconv.Itoa(user.ID)) {
		return fmt.Sprint(fileCache.Get(USER_TOKEN + strconv.Itoa(user.ID)))
	} else {
		return ""
	}
}
