package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type SystemChip struct{}

// Group 系统筹码分组
func (g *SystemChip) Group(r *gin.RouterGroup) {
	group := r.Group("system_chip")
	{
		group.POST("list", admin.Entry.SystemChip.List) // 列表
	}
	{
		group.GET("getBox", admin.Entry.SystemChip.GetBox) // 选择框
	}
}
