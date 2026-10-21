package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type Currency struct{}

// Group 币种分组
func (g *Currency) Group(r *gin.RouterGroup) {
	group := r.Group("currency")
	{
		group.POST("list", admin.Entry.Currency.List)                                       // 列表
		group.POST("add", admin.Entry.Currency.Add)                                         // 创建
		group.POST("del", admin.Entry.Currency.Del)                                         // 删除
		group.POST("setExchangeRate", admin.Entry.Currency.SetExchangeRate)                 // 设置币种汇率
		group.POST("exchangeRateHistoryList", admin.Entry.Currency.ExchangeRateHistoryList) // 币种汇率历史记录
	}
	{
		group.GET("detail/:id", admin.Entry.Currency.Detail) // 详情
	}
}
