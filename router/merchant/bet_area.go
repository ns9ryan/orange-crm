package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type BetArea struct{}

// Group 下注区域分组
func (g *BetArea) Group(r *gin.RouterGroup) {
	group := r.Group("bet_area")
	{
		group.POST("list", merchant.Entry.BetArea.List)                               // 列表
		group.POST("getBoxByCategorySub", merchant.Entry.BetArea.GetBoxByCategorySub) // 根据子分类获取下拉框数据
	}
}
