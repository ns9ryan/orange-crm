package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type GameCategory struct{}

// Group 游戏分类分组
func (g *GameCategory) Group(r *gin.RouterGroup) {
	group := r.Group("game_category")
	{
		group.POST("list", merchant.Entry.GameCategory.List) // 列表
	}
	{
		group.GET("getBox", merchant.Entry.GameCategory.GetBox) // 选择框
	}
}
