package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type Menu struct{}

// Group 菜单分组
func (g *Menu) Group(r *gin.RouterGroup) {
	group := r.Group("menu")
	{
		group.POST("list", admin.Entry.Menu.List) // 列表
		group.POST("add", admin.Entry.Menu.Add)   // 创建
		group.POST("edit", admin.Entry.Menu.Edit) // 编辑
		group.POST("del", admin.Entry.Menu.Del)   // 删除
	}
	{
		group.GET("detail/:id", admin.Entry.Menu.Detail) // 详情
	}
}
