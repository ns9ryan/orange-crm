package merchant

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/merchant"
)

type Merchant struct{}

// Group 商户分组
func (g *Merchant) Group(loginGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	loginGroup = loginGroup.Group("merchant")

	// 需要登录并且需要权限的路由分组
	privateGroup = privateGroup.Group("merchant")
	{
		privateGroup.POST("list", merchant.Entry.Merchant.List)                             // 列表
		privateGroup.POST("addChildren", merchant.Entry.Merchant.AddChildren)               // 创建子商户
		privateGroup.POST("edit", merchant.Entry.Merchant.Edit)                             // 编辑
		privateGroup.POST("del", merchant.Entry.Merchant.Del)                               // 删除
		privateGroup.POST("checkUsername", merchant.Entry.Merchant.CheckUsername)           // 检查账号是否可用
		loginGroup.POST("editCurrentPassword", merchant.Entry.Merchant.EditCurrentPassword) // 编辑当前登录的商户密码
		loginGroup.POST("resetPassword", merchant.Entry.Merchant.ResetPassword)             // 重置密码
	}
	{
		privateGroup.GET("getApiKey/:id", merchant.Entry.Merchant.GetApiKey)                           // 获取api相关密钥
		privateGroup.GET("detail/:id", merchant.Entry.Merchant.Detail)                                 // 详情
		privateGroup.GET("detailByMerchantKey/:username", merchant.Entry.Merchant.DetailByMerchantKey) // 通过key获取详情
		privateGroup.GET("regenerateAesKeySalt/:id", merchant.Entry.Merchant.RegenerateAesKeySalt)     // 重新生成AesKeySalt
		privateGroup.GET("regenerateMd5Salt/:id", merchant.Entry.Merchant.RegenerateMd5Salt)           // 重新生成Md5Salt
		privateGroup.GET("regenerateTokenKey/:id", merchant.Entry.Merchant.RegenerateTokenKey)         // 重新生成商户tokenKey
		privateGroup.GET("getSubBox", merchant.Entry.Merchant.GetSubBox)                               // 子商户下拉框数据
	}
}
