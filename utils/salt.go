package utils

import (
	"math/rand"
	"orange-crm/global"
	"strings"
)

// GenerateRandomString 随机字符串生成函数
func GenerateRandomString(r *rand.Rand, length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(global.Charset[rand.Intn(len(global.Charset))])
	}
	return sb.String()
}
