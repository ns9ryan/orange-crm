package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type Player struct{}

// Group 玩家分组
func (g *Player) Group(r *gin.RouterGroup) {
	group := r.Group("player")
	{
		group.POST("list", merchant.Entry.Player.List)                   // 列表
		group.POST("add", merchant.Entry.Player.Add)                     // 创建
		group.POST("edit", merchant.Entry.Player.Edit)                   // 编辑
		group.POST("resetPassword", merchant.Entry.Player.ResetPassword) // 重置密码
	}
	{
		group.GET("detail/:id", merchant.Entry.Player.Detail) // 详情
	}
}
