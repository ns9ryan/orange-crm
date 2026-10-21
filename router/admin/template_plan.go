package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type TemplatePlan struct{}

// Group 限红方案分组
func (g *TemplatePlan) Group(r *gin.RouterGroup) {
	group := r.Group("template_plan")
	{
		group.POST("list", admin.Entry.TemplatePlan.List)                               // 列表
		group.POST("add", admin.Entry.TemplatePlan.Add)                                 // 创建
		group.POST("edit", admin.Entry.TemplatePlan.Edit)                               // 编辑
		group.POST("del", admin.Entry.TemplatePlan.Del)                                 // 删除
		group.POST("getBoxByCategorySub", admin.Entry.TemplatePlan.GetBoxByCategorySub) // 根据子分类获取下拉框数据
	}
	{
		group.GET("detail/:id", admin.Entry.TemplatePlan.Detail) // 详情
	}
}
