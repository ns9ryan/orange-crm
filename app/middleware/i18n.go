package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"orange-crm/global"
)

// UseI18n 国际化语言
func UseI18n() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Request.Header.Get("x-lang")
		if len(lang) == 0 {
			lang = global.CONFIG.Server.Lang
		}
		global.I18N.Context = gi18n.WithLanguage(c, lang)
		c.Next()
	}
}
