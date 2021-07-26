package util

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenToken(s string) string {
	now := time.Now()
	rs := RandStringRunes(20) + s + now.String()
	h := sha1.New()
	h.Write([]byte(rs))
	bs := h.Sum(nil)
	token := fmt.Sprintf("%x", bs)
	return token
}
