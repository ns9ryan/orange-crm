package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type SystemCurrency struct{}

// Group 系统币种分组
func (g *SystemCurrency) Group(r *gin.RouterGroup) {
	group := r.Group("system_currency")
	{
		group.POST("list", admin.Entry.SystemCurrency.List)                                    // 列表
		group.GET("getSystemCurrencyNoUse", admin.Entry.SystemCurrency.GetSystemCurrencyNoUse) // 获取未使用的系统币种
	}
}
