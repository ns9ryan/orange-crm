package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type BetArea struct{}

// Group 下注区域分组
func (g *BetArea) Group(r *gin.RouterGroup) {
	group := r.Group("bet_area")
	{
		group.POST("list", admin.Entry.BetArea.List)                               // 列表
		group.POST("getBoxByCategorySub", admin.Entry.BetArea.GetBoxByCategorySub) // 根据子分类获取下拉框数据
	}
}
