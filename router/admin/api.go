package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type Api struct{}

// Group api管理分组
func (g *Api) Group(r *gin.RouterGroup) {
	group := r.Group("api")
	{
		group.POST("list", admin.Entry.Api.List) // 列表
		group.POST("add", admin.Entry.Api.Add)   // 创建
		group.POST("edit", admin.Entry.Api.Edit) // 编辑
		group.POST("del", admin.Entry.Api.Del)   // 删除
	}
	{
		group.GET("detail/:id", admin.Entry.Api.Detail)             // 详情
		group.GET("getApiGroupBox", admin.Entry.Api.GetApiGroupBox) // 获取api分组选择框
	}
}
