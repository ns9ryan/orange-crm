// =================================================================================
// 路由-控端账号
// =================================================================================

package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type ControlAccount struct{}

// Group 控端账号分组
func (g *ControlAccount) Group(r *gin.RouterGroup) {
	group := r.Group("control_account")
	{
		group.POST("list", admin.Entry.ControlAccount.List)                   // 列表
		group.POST("add", admin.Entry.ControlAccount.Add)                     // 创建
		group.POST("edit", admin.Entry.ControlAccount.Edit)                   // 编辑
		group.POST("del", admin.Entry.ControlAccount.Del)                     // 删除
		group.POST("resetPassword", admin.Entry.ControlAccount.ResetPassword) // 重置密码
	}
	{
		group.GET("detail/:id", admin.Entry.ControlAccount.Detail)       // 详情
		group.GET("getLevelBox", admin.Entry.ControlAccount.GetLevelBox) // 获取控端账号等级选择框
		group.GET("getBox", admin.Entry.ControlAccount.GetBox)           // 获取控端账号下拉框数据
	}
}
