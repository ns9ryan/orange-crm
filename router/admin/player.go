package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type Player struct{}

// Group 玩家分组
func (g *Player) Group(r *gin.RouterGroup) {
	group := r.Group("player")
	{
		group.POST("list", admin.Entry.Player.List)                   // 列表
		group.POST("add", admin.Entry.Player.Add)                     // 创建
		group.POST("edit", admin.Entry.Player.Edit)                   // 编辑
		group.POST("del", admin.Entry.Player.Del)                     // 删除
		group.POST("resetPassword", admin.Entry.Player.ResetPassword) // 重置密码
	}
	{
		group.GET("detail/:id", admin.Entry.Player.Detail) // 详情
	}
}
