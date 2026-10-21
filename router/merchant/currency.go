package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type Currency struct{}

// Group 币种分组
func (g *Currency) Group(r *gin.RouterGroup) {
	group := r.Group("currency")
	{
		group.POST("list", merchant.Entry.Currency.List) // 列表
	}
	{
		group.GET("detail/:id", merchant.Entry.Currency.Detail) // 详情
	}
}
