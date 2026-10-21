package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type Dealer struct{}

// Group 荷官分组
func (g *Dealer) Group(r *gin.RouterGroup) {
	group := r.Group("dealer")
	{
		group.POST("uploadAvatar", admin.Entry.Dealer.UploadAvatar) // 上传头像
		group.POST("list", admin.Entry.Dealer.List)                 // 列表
		group.POST("add", admin.Entry.Dealer.Add)                   // 创建
		group.POST("edit", admin.Entry.Dealer.Edit)                 // 编辑
		group.POST("del", admin.Entry.Dealer.Del)                   // 删除
	}
	{
		group.GET("detail/:id", admin.Entry.Dealer.Detail)                             // 详情
		group.GET("generateEmployeeNumber", admin.Entry.Dealer.GenerateEmployeeNumber) // 生成工号
	}
}
