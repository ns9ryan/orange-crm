// =================================================================================
// 路由-会员交易流水
// =================================================================================

package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type PlayerTransaction struct{}

// Group 会员交易流水分组
func (g *PlayerTransaction) Group(r *gin.RouterGroup) {
	group := r.Group("player_transaction")
	{
		group.POST("creditAndDebitList", admin.Entry.PlayerTransaction.CreditAndDebitList) // 上分/下分记录列表
	}
}
