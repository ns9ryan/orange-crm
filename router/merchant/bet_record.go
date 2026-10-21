package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type BetRecord struct{}

// Group 投注明细分组
func (g *BetRecord) Group(r *gin.RouterGroup) {
	group := r.Group("bet_record")
	{
		group.POST("list", merchant.Entry.BetRecord.List) // 列表
	}
}
