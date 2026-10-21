package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
	"orange-crm/app/service"
	"orange-crm/global"
	"orange-crm/utils"
	"strconv"
	"strings"
)

// CasbinMiddleware casbin中间件
func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("or-token")
		if len(token) == 0 {
			utils.Response.FailUnauthorized(c)
			c.Abort()
			return
		}

		claims, err := service.ParseToken(token)
		if err != nil {
			global.LOG.Error("解析token失败：", zap.Error(err))
			utils.Response.FailUnauthorized(c)
			c.Abort()
			return
		}

		// 超级管理员账号不需要验证权限
		if claims.IsSuper == 1 {
			c.Next()
			return
		}

		// TODO 商户主账号暂时不需要验证权限
		//if claims.IsMain == 1 {
		//	c.Next()
		//	return
		//}

		//sub := "admin_1"
		sub := ""
		// 根据身份判断sub
		if claims.Identity == global.AdminIdentity { // 管理员身份
			sub = "admin_" + gconv.String(claims.ID)
		} else if claims.Identity == global.MerchantIdentity { // 商户身份
			sub = "merchant_" + strconv.Itoa(int(claims.ID))
		}

		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := strings.ToUpper(c.Request.Method)
		//id := "1"

		// 获取casbin服务	sub="admin_1" obj="/role/save" act="POST" id="1"
		ok, err := global.CASBIN.Enforce(sub, obj, act)
		if err != nil {
			global.LOG.Error("权限验证出错：", zap.Error(err))
			utils.Response.FailUnauthorized(c)
			c.Abort()
			return
		}

		// 没有权限
		if !ok {
			utils.Response.FailForbidden(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
