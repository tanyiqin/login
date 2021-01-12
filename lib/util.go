package lib

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// 生成32位MD5
func MD5(text string, len int) []byte{
	ctx := md5.New()
	ctx.Write([]byte(text))
	Salt := GetRandomString(32)
	return hex.EncodeToString(ctx.Sum(Salt))
}

//生成随机字符串
func GetRandomString(length int) []byte{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}