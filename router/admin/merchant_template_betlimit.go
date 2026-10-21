package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type MerchantTemplateBetlimit struct{}

// Group 商户限红模板分组
func (g *MerchantTemplateBetlimit) Group(r *gin.RouterGroup) {
	group := r.Group("merchant_template_betlimit")
	{
		group.POST("list", admin.Entry.MerchantTemplateBetlimit.List)                                         // 列表
		group.POST("add", admin.Entry.MerchantTemplateBetlimit.Add)                                           // 创建
		group.POST("edit", admin.Entry.MerchantTemplateBetlimit.Edit)                                         // 编辑
		group.POST("del", admin.Entry.MerchantTemplateBetlimit.Del)                                           // 删除
		group.POST("getTemplateBoxByMerchant", admin.Entry.MerchantTemplateBetlimit.GetTemplateBoxByMerchant) // 获取商户限红模板下拉框数据
	}
	{
		group.GET("detail/:id", admin.Entry.MerchantTemplateBetlimit.Detail) // 详情
	}
}
