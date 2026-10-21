package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type Admin struct{}

// Group 管理员分组
func (g *Admin) Group(loginGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	loginGroup = loginGroup.Group("admin")

	// 需要登录并且需要权限的路由分组
	privateGroup = privateGroup.Group("admin")

	{
		privateGroup.POST("list", admin.Entry.Admin.List)                             // 列表
		privateGroup.POST("add", admin.Entry.Admin.Add)                               // 创建
		privateGroup.POST("edit", admin.Entry.Admin.Edit)                             // 编辑
		privateGroup.POST("del", admin.Entry.Admin.Del)                               // 删除
		loginGroup.POST("editCurrentPassword", admin.Entry.Admin.EditCurrentPassword) // 编辑当前登录的管理员密码
		privateGroup.POST("resetPassword", admin.Entry.Admin.ResetPassword)           // 重置密码
		loginGroup.POST("getAdminInfo", admin.Entry.Admin.GetAdminInfo)               // 获取管理员信息
		loginGroup.POST("logoutAdmin", admin.Entry.Admin.LogoutAdmin)                 // 管理员登出
	}
	{
		privateGroup.GET("detail/:id", admin.Entry.Admin.Detail) // 详情
	}
}
