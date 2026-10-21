// =================================================================================
// 路由-会员交易流水
// =================================================================================

package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type PlayerTransaction struct{}

// Group 会员交易流水分组
func (g *PlayerTransaction) Group(r *gin.RouterGroup) {
	group := r.Group("player_transaction")
	{
		group.POST("creditAndDebitList", merchant.Entry.PlayerTransaction.CreditAndDebitList) // 上分/下分记录列表
	}
}
