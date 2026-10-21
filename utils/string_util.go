package utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

// TruncateString 截断字符串到指定长度，超过部分追加标记
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "...【已截断】"
}

// GetCtxString 从 Gin Context 中读取指定 key 的 string 值
func GetCtxString(c *gin.Context, key string) string {
	if val, exists := c.Get(key); exists {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

// GetCtxDurationMillis 获取 Duration 类型上下文值并转换为毫秒（uint）
func GetCtxDurationMillis(c *gin.Context, key string) uint {
	if val, exists := c.Get(key); exists {
		if d, ok := val.(time.Duration); ok {
			return uint(d.Milliseconds())
		}
	}
	return 0
}
