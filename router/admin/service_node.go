package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type ServiceNode struct{}

// Group 服务节点分组
func (g *ServiceNode) Group(r *gin.RouterGroup) {
	group := r.Group("service_node")
	{
		group.POST("list", admin.Entry.ServiceNode.List) // 列表
		group.POST("add", admin.Entry.ServiceNode.Add)   // 创建
		group.POST("edit", admin.Entry.ServiceNode.Edit) // 编辑
		group.POST("del", admin.Entry.ServiceNode.Del)   // 删除
	}
	{
		group.GET("detail/:id", admin.Entry.ServiceNode.Detail) // 详情
	}
}
