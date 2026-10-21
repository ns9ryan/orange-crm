package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type MerchantDayaccount struct{}

// Group 商户输赢报表分组
func (g *MerchantDayaccount) Group(r *gin.RouterGroup) {
	group := r.Group("merchant_dayaccount")
	{
		group.POST("list", admin.Entry.MerchantDayaccount.List)       // 获取主商户输赢报表
		group.POST("subList", admin.Entry.MerchantDayaccount.SubList) // 获取主商户下子商户的输赢报表
	}
}
