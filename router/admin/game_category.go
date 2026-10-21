package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type GameCategory struct{}

// Group 游戏分类分组
func (g *GameCategory) Group(r *gin.RouterGroup) {
	group := r.Group("game_category")
	{
		group.POST("list", admin.Entry.GameCategory.List) // 列表
	}
	{
		group.GET("getBox", admin.Entry.GameCategory.GetBox) // 选择框
	}
}
