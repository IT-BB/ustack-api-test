package framework

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

// RandomName 生成带前缀的随机名称
func RandomName(prefix string) string {
	return fmt.Sprintf("%s-%d-%s", prefix, time.Now().Unix(), randString(6))
}

func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Ptr 返回值的指针，用于构建参数
func Ptr[T any](v T) *T {
	return &v
}

// EncodeTags 将 key:value map 转换为 API 要求的 Base64 格式数组
// 格式: Base64(Key):Base64(Value)
func EncodeTags(tags map[string]string) []string {
	var res []string
	for k, v := range tags {
		bk := base64.StdEncoding.EncodeToString([]byte(k))
		bv := base64.StdEncoding.EncodeToString([]byte(v))
		res = append(res, fmt.Sprintf("%s:%s", bk, bv))
	}
	return res
}
