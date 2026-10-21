package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type SystemNotice struct{}

// Group 系统公告分组
func (g *SystemNotice) Group(r *gin.RouterGroup) {
	group := r.Group("system_notice")
	{
		group.POST("list", admin.Entry.SystemNotice.List) // 列表
		group.POST("add", admin.Entry.SystemNotice.Add)   // 创建
		group.POST("edit", admin.Entry.SystemNotice.Edit) // 编辑
		group.POST("del", admin.Entry.SystemNotice.Del)   // 删除
	}
	{
		group.GET("detail/:id", admin.Entry.SystemNotice.Detail) // 详情
	}
}
