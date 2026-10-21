package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type TemplatePlan struct{}

// Group 限红方案分组
func (g *TemplatePlan) Group(r *gin.RouterGroup) {
	group := r.Group("template_plan")
	{
		group.POST("list", merchant.Entry.TemplatePlan.List)                               // 列表
		group.POST("getBoxByCategorySub", merchant.Entry.TemplatePlan.GetBoxByCategorySub) // 根据子分类获取下拉框数据
	}
	{
		group.GET("detail/:id", merchant.Entry.TemplatePlan.Detail) // 详情
	}
}
