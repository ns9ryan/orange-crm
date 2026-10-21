package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type LogModule struct{}

// Group 日志分组
func (g *LogModule) Group(r *gin.RouterGroup) {
	group := r.Group("log_module")
	{
		group.POST("adminLoginLogList", admin.Entry.LogModule.AdminLoginLogList)                      // 管理员登录日志列表
		group.POST("merchantLoginLogList", admin.Entry.LogModule.MerchantLoginLogList)                // 商户登录日志列表
		group.POST("adminOperationLogList", admin.Entry.LogModule.AdminOperationLogList)              // 管理员操作日志列表
		group.GET("getAdminOperationLogEventBox", admin.Entry.LogModule.GetAdminOperationLogEventBox) // 管理员操作日志事件下拉框
		group.POST("playerLoginLogList", admin.Entry.LogModule.PlayerLoginLogList)                    // 会员登录日志列表
		group.POST("controlLoginLogList", admin.Entry.LogModule.ControlLoginLogList)                  // 控端登录日志列表
	}
}
