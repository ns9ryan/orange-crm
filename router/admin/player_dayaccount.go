package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type PlayerDayaccount struct{}

// Group 玩家输赢表分组
func (g *PlayerDayaccount) Group(r *gin.RouterGroup) {
	group := r.Group("player_dayaccount")
	{
		group.POST("list", admin.Entry.PlayerDayaccount.List) // 列表
	}
}
