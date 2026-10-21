package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type LogModule struct{}

// Group 日志分组
func (g *LogModule) Group(r *gin.RouterGroup) {
	group := r.Group("log_module")
	{
		group.POST("merchantLoginLogList", merchant.Entry.LogModule.MerchantLoginLogList)                      // 商户登录日志列表
		group.POST("merchantOperationLogList", merchant.Entry.LogModule.MerchantOperationLogList)              // 商户操作日志列表
		group.GET("getMerchantOperationLogEventBox", merchant.Entry.LogModule.GetMerchantOperationLogEventBox) // 商户操作日志事件下拉框
		group.POST("playerLoginLogList", merchant.Entry.LogModule.PlayerLoginLogList)                          // 会员登录日志列表
	}
}
