package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type GameAnswer struct{}

// Group 游戏开答分组
func (g *GameAnswer) Group(r *gin.RouterGroup) {
	group := r.Group("game_answer")
	{
		group.POST("list", merchant.Entry.GameAnswer.List) // 列表
	}
}
