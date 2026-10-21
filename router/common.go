package router

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler"
)

type Common struct{}

func (g *Common) Group(r *gin.RouterGroup) {
	group := r.Group("")
	common := handler.Common{}
	{
		group.GET("getLangBox", common.GetLangBox)           // 语言下拉框
		group.GET("getSystemConfig", common.GetSystemConfig) // 获取系统配置数据
		group.GET("gameServiceName", common.GameServiceName) // 获取游戏服务名称
	}
	{
		group.POST("loginAdmin", common.LoginAdmin)       // 管理员登录
		group.POST("loginMerchant", common.LoginMerchant) // 商户登录
		group.POST("imgUpload", common.UploadImg)         // 上传图片
		group.POST("chunkUpload", common.UploadChunk)     // 通用上传切片
		group.POST("loginPlayer", common.LoginPlayer)     // 会员登录-测试
	}
}
