package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type MerchantDayaccount struct{}

// Group 商户输赢报表分组
func (g *MerchantDayaccount) Group(r *gin.RouterGroup) {
	group := r.Group("merchant_dayaccount")
	{
		//
		group.POST("list", merchant.Entry.MerchantDayaccount.List) // 获取商户自身的输赢报表
		//
		group.POST("subList", merchant.Entry.MerchantDayaccount.SubList) // 获取主商户下子商户的输赢报表（仅主商户可用）
	}
}
