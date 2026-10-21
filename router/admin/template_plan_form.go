package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type TemplatePlanForm struct{}

// Group 限红模板分组
func (g *TemplatePlanForm) Group(r *gin.RouterGroup) {
	group := r.Group("template_plan_form")
	{
		group.POST("list", admin.Entry.TemplatePlanForm.List) // 列表
		group.POST("add", admin.Entry.TemplatePlanForm.Add)   // 创建
		group.POST("edit", admin.Entry.TemplatePlanForm.Edit) // 编辑
		group.POST("del", admin.Entry.TemplatePlanForm.Del)   // 删除
	}
	{
		group.GET("detail/:id", admin.Entry.TemplatePlanForm.Detail) // 详情
		group.GET("getBox", admin.Entry.TemplatePlanForm.GetBox)     // 限红模板下拉框数据
	}
}
