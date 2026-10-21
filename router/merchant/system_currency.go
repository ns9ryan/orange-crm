package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type SystemCurrency struct{}

// Group 系统币种分组
func (g *SystemCurrency) Group(r *gin.RouterGroup) {
	group := r.Group("system_currency")
	{
		group.POST("list", merchant.Entry.SystemCurrency.List)                                    // 列表
		group.GET("getSystemCurrencyNoUse", merchant.Entry.SystemCurrency.GetSystemCurrencyNoUse) // 获取未使用的系统币种
	}
}
