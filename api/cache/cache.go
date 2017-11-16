package cache

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var fileCache cache.Cache

func init() {
	fmt.Println("缓存初始化")
	// bm, _ := cache.NewCache("file", `{"CachePath":"./file_cache","FileSuffix":".cache","DirectoryLevel":"1","EmbedExpiry":12000}`)
	bm, _ := cache.NewCache("redis", `{"key":"redis","conn":":6379","dbNum":"0","password":""}`)
	fileCache = bm
}

//生成随机字符串
func GenerateToken() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetValue 获取缓存key对应的值
func getValue(key string) string {
	if fileCache.IsExist(key) {
		return string(fileCache.Get(key).([]byte))
	}
	return ""
}
