package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type GameTable struct{}

// Group 游戏桌台分组
func (g *GameTable) Group(r *gin.RouterGroup) {
	group := r.Group("game_table")
	{
		group.POST("list", merchant.Entry.GameTable.List) // 列表
	}
	{
		group.GET("detail/:id", merchant.Entry.GameTable.Detail)               // 详情
		group.GET("getBoxWithGroup", merchant.Entry.GameTable.GetBoxWithGroup) // 获取游戏桌台,分组下拉框
		group.GET("getBox", merchant.Entry.GameTable.GetBox)                   // 游戏桌台下拉框
	}
}
