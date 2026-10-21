package router

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler"
)

type DataInit struct{}

func (g *DataInit) Group(r *gin.RouterGroup) {
	group := r.Group("data_init")
	dataInit := handler.DataInit{}
	{
		group.POST("addMerchants", dataInit.AddMerchants) // 批量添加商户
		group.POST("addPlayers", dataInit.AddPlayers)     // 批量添加玩家
		group.POST("deposit", dataInit.Deposit)           // 批量存款上分

		group.GET("editLoginLog", dataInit.EditLoginLog) // 更新登录日志

		group.GET("editPlayerTransaction", dataInit.EditPlayerTransaction) // 更新玩家流水表的商户字段，后面增加的字段没有赋值
		group.GET("editPlayerWallet", dataInit.EditPlayerWallet)           // 更新玩家钱包表的商户字段，后面增加的字段没有赋值

		group.GET("editMerchantTips", dataInit.EditMerchantTips) // 更新商户的打赏表
		group.POST("makeToken", dataInit.MakeToken)              // 生成商户token
	}
}
