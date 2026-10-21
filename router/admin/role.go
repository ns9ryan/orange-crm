package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type Role struct{}

// Group 角色分组
func (g *Role) Group(r *gin.RouterGroup) {
	group := r.Group("role")
	{
		group.POST("list", admin.Entry.Role.List)                           // 列表
		group.POST("add", admin.Entry.Role.Add)                             // 创建
		group.POST("edit", admin.Entry.Role.Edit)                           // 编辑
		group.POST("del", admin.Entry.Role.Del)                             // 删除
		group.POST("setRoleMenus", admin.Entry.Role.SetRoleMenus)           // 设置角色菜单
		group.POST("setRoleApis", admin.Entry.Role.SetRoleApis)             // 设置角色Apis
		group.POST("getMenuListByRole", admin.Entry.Role.GetMenuListByRole) // 根据角色获取菜单列表
	}
	{
		group.GET("detail/:id", admin.Entry.Role.Detail) // 详情
	}
}
