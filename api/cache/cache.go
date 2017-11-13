package cache

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"hkllzh.com/easy-bill/api/models"
	"strconv"
	"fmt"
	"time"
	"math/rand"
	"github.com/astaxie/beego/logs"
)

var fileCache cache.Cache

const UserToken = "user_token_"

func init() {
	fmt.Println("缓存初始化")
	// bm, _ := cache.NewCache("file", `{"CachePath":"./file_cache","FileSuffix":".cache","DirectoryLevel":"1","EmbedExpiry":12000}`)
	bm, _ := cache.NewCache("redis", `{"key":"redis","conn":":6379","dbNum":"0","password":""}`)
	fileCache = bm
}

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

//生成随机字符串
func GetToken() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
