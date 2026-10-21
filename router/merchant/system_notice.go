package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type SystemNotice struct{}

// Group 系统公告分组
func (g *SystemNotice) Group(r *gin.RouterGroup) {
	group := r.Group("system_notice")
	{
		group.POST("listForNew", merchant.Entry.SystemNotice.ListForNew)   // 商户新消息
		group.POST("listForSelf", merchant.Entry.SystemNotice.ListForSelf) // 商户收到的系统公共列表
		group.POST("markRead", merchant.Entry.SystemNotice.MarkRead)       // 标记已读
	}
}
