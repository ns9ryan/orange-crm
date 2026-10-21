// =================================================================================
// 路由-商户
// =================================================================================

package admin

import (
	"github.com/gin-gonic/gin"
	"orange-crm/app/handler/admin"
)

type Merchant struct{}

// Group 商户分组
func (g *Merchant) Group(r *gin.RouterGroup) {
	group := r.Group("merchant")
	{
		group.POST("list", admin.Entry.Merchant.List)                   // 列表
		group.POST("add", admin.Entry.Merchant.Add)                     // 创建
		group.POST("addChildren", admin.Entry.Merchant.AddChildren)     // 创建子商户
		group.POST("edit", admin.Entry.Merchant.Edit)                   // 编辑
		group.POST("del", admin.Entry.Merchant.Del)                     // 删除
		group.POST("checkUsername", admin.Entry.Merchant.CheckUsername) // 检查账号是否可用
		group.POST("resetPassword", admin.Entry.Merchant.ResetPassword) // 重置密码
	}
	{
		group.GET("getApiKey/:id", admin.Entry.Merchant.GetApiKey)                               // 获取api相关密钥
		group.GET("detail/:id", admin.Entry.Merchant.Detail)                                     // 详情
		group.GET("detailByMerchantKey/:merchant_key", admin.Entry.Merchant.DetailByMerchantKey) // 通过key获取详情
		group.GET("regenerateAesKeySalt/:id", admin.Entry.Merchant.RegenerateAesKeySalt)         // 重新生成AesKeySalt
		group.GET("regenerateMd5Salt/:id", admin.Entry.Merchant.RegenerateMd5Salt)               // 重新生成Md5Salt
		group.GET("regenerateTokenKey/:id", admin.Entry.Merchant.RegenerateTokenKey)             // 重新生成商户tokenKey
		group.GET("getBox", admin.Entry.Merchant.GetBox)                                         // 商户下拉框数据
		group.GET("getParentBox", admin.Entry.Merchant.GetParentBox)                             // 主商户下拉框数据
		group.GET("getSubBox/:merchant_id", admin.Entry.Merchant.GetSubBox)                      // 子商户下拉框数据
	}

}
