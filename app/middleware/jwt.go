package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"orange-crm/app/service"
	"orange-crm/global"
	"orange-crm/utils"
)

// CheckLogin 是否登录
func CheckLogin() gin.HandlerFunc {
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

		//// 判断Redis中是否存在token
		//jsonData, err := global.Redis.Get(context.Background(), token).Result()
		//if err != nil {
		//	global.LOG.Error("获取redis登录信息失败：", zap.Error(err))
		//	utils.Response.FailUnauthorized(c)
		//	c.Abort()
		//	return
		//}
		//
		//claims := response.LoginResponse{}
		//if err = json.Unmarshal([]byte(jsonData), &claims); err != nil {
		//	utils.Response.FailWithDetail(c, err.Error())
		//	c.Abort()
		//	return
		//}

		c.Set("userinfo", claims)

		c.Next()
	}
}

// CheckAdmin 是否管理员
func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取登录信息
		userInfo, err := service.GetAdminInfo(c)
		if err != nil {
			utils.Response.FailUnauthorized(c)
			c.Abort()
			return
		}

		if userInfo.Identity != global.AdminIdentity {
			utils.Response.FailIdentity(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
