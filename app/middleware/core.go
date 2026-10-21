package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)

		// 允许所有的请求头
		//c.Header("Access-Control-Allow-Headers", "*")

		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization,X-Lang ,Token,X-Token,Or-Token,ContentType")

		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// StaticServerWithCORS 全局静态资源带 CORS 中间件
func StaticServerWithCORS(relativePath string, root string) gin.HandlerFunc {
	fs := http.Dir(root)
	fileServer := http.FileServer(fs)

	return func(c *gin.Context) {
		// 检查请求是不是访问静态资源路径
		if strings.HasPrefix(c.Request.URL.Path, relativePath) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
			c.Header("Access-Control-Allow-Credentials", "true")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}

			// 去除 URL 的前缀再交给 FileServer
			http.StripPrefix(relativePath, fileServer).ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}

		// 如果不是静态资源，就交给后面路由继续处理
		c.Next()
	}
}
