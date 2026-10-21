package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type GameTable struct{}

// Group 游戏桌台分组
func (g *GameTable) Group(r *gin.RouterGroup) {
	group := r.Group("game_table")
	{
		group.POST("list", admin.Entry.GameTable.List) // 列表
		group.POST("add", admin.Entry.GameTable.Add)   // 创建
		group.POST("edit", admin.Entry.GameTable.Edit) // 编辑
		group.POST("del", admin.Entry.GameTable.Del)   // 删除
	}
	{
		group.GET("detail/:id", admin.Entry.GameTable.Detail)               // 详情
		group.GET("getBoxWithGroup", admin.Entry.GameTable.GetBoxWithGroup) // 获取游戏桌台,分组下拉框
		group.GET("getBox", admin.Entry.GameTable.GetBox)                   // 游戏桌台下拉框
	}
}
