package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type MerchantNotice struct{}

// Group 商户公告分组
func (g *MerchantNotice) Group(r *gin.RouterGroup) {
	group := r.Group("merchant_notice")
	{
		group.POST("list", merchant.Entry.MerchantNotice.List) // 列表
		group.POST("add", merchant.Entry.MerchantNotice.Add)   // 创建
		group.POST("edit", merchant.Entry.MerchantNotice.Edit) // 编辑
		group.POST("del", merchant.Entry.MerchantNotice.Del)   // 删除
	}
	{
		group.GET("detail/:id", merchant.Entry.MerchantNotice.Detail) // 详情
	}
}
